package ethbase

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	errors "github.com/pkg/errors"

	relaynetwork "github.com/planet-ethereum/relay-network"
	wallet "github.com/planet-ethereum/relay-network/wallet"
)

type Subscription struct {
	EventId   [32]byte
	Emitter   common.Address
	EventName [32]byte
	Account   common.Address
	Method    [4]byte
}

type EthbaseMessage struct {
	EventId [32]byte `json:"eventId"`
	Data    []byte   `json:"data"` // Event parameters
	Proof   LogProof `json:"proof"`
}

type Ethbase struct {
	address       common.Address
	client        *ethclient.Client
	registry      *Registry
	subscriptions []Subscription
}

func NewEthbase(client *ethclient.Client, address common.Address) (*Ethbase, error) {
	e := new(Ethbase)

	e.client = client
	e.address = address

	var err error
	e.registry, err = NewRegistry(e.address, e.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create registry instance")
	}

	return e, nil
}

func (e *Ethbase) Initialize() error {
	subscriptions, err := e.GetSubscriptions()
	if err != nil {
		return err
	}

	e.subscriptions = subscriptions

	return nil
}

func (e *Ethbase) GetSubscriptions() ([]Subscription, error) {
	logs, err := getPastLogs(e.client, e.address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch past logs")
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(RegistryABI)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct abi from json")
	}

	res := make([]Subscription, 0, len(logs))
	for _, l := range logs {
		sub := Subscription{}
		if err = contractAbi.Unpack(&sub, "Subscribed", l.Data); err != nil {
			return nil, errors.Wrap(err, "failed to unpack subscribed log")
		}

		res = append(res, sub)
	}

	return res, nil
}

func (e *Ethbase) LogToMessage(ctx context.Context, l types.Log) (*relaynetwork.Message, error) {
	for _, sub := range e.subscriptions {
		log.Printf(
			"EventId: %v\tEmitter: %s\tAccount: %s\tData: %v\n",
			hex.EncodeToString(sub.EventId[:]),
			sub.Emitter.Hex(),
			sub.Account.Hex(),
			hex.EncodeToString(sub.Method[:]),
		)

		proof, err := GenerateLogProof(ctx, e.client, l)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate proof")
		}

		em := &EthbaseMessage{
			Data:    l.Data,
			EventId: sub.EventId,
			Proof:   *proof,
		}
		raw, err := json.Marshal(em)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshalized ethbase message")
		}

		if sub.Emitter == l.Address {
			m := &relaynetwork.Message{
				Kind: "ethbase",
				To:   sub.Account.Hex(),
				Data: raw,
			}

			return m, nil
		}
	}

	return nil, errors.New("no subscription found")
}

func (e *Ethbase) SubmitTx(ctx context.Context, wal *wallet.HDWallet, m relaynetwork.Message) error {
	if m.Kind != "ethbase" {
		return errors.New("tx has wrong kind")
	}

	em := EthbaseMessage{}
	if err := json.Unmarshal(m.Data, &em); err != nil {
		return err
	}

	fromAddress, err := wal.Address()
	if err != nil {
		return err
	}

	nonce, err := e.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	gasLimit := uint64(300000)

	auth, err := wal.NewKeyedTransactor()
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	subscriber := common.HexToAddress(m.To)

	tx, err := e.registry.SubmitLog(
		auth,
		em.Proof.Value,
		em.Proof.Proof,
		em.Proof.Key,
		big.NewInt(int64(em.Proof.LogIndex)),
		em.Proof.Header,
		subscriber,
		em.EventId,
	)

	log.Printf("TX sent: %s\n", tx.Hash().Hex())

	return err
}

// Listen subscribes to incoming logs.
func (e *Ethbase) Listen(ctx context.Context, msgCh chan []byte, errCh chan error, emitters []common.Address) error {
	query := ethereum.FilterQuery{Addresses: emitters}

	logsCh := make(chan types.Log)
	sub, err := e.client.SubscribeFilterLogs(ctx, query, logsCh)
	if err != nil {
		return err
	}

	log.Println("Subscribed to logs")
	go e.listen(ctx, sub, logsCh, msgCh, errCh)

	return nil
}

func (e *Ethbase) listen(ctx context.Context, sub ethereum.Subscription, logsCh chan types.Log, msgCh chan []byte, errCh chan error) {
	for {
		select {
		case err := <-sub.Err():
			sub.Unsubscribe()
			errCh <- err
			return
		case vLog := <-logsCh:
			m, err := e.LogToMessage(ctx, vLog)
			if err != nil {
				log.Printf("failed to convert log to msg: %v\n", err)
				continue
			}

			marshalized, err := json.Marshal(m)
			if err != nil {
				errCh <- errors.Wrap(err, "failed to marshalize message")
			}

			msgCh <- marshalized
		}
	}
}

func getPastLogs(client *ethclient.Client, address common.Address) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			address,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

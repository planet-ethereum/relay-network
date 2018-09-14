package ethbase

import (
	"context"
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

func (e *Ethbase) LogToMessage(l types.Log) (relaynetwork.Message, bool) {
	for _, sub := range e.subscriptions {
		log.Printf("EventId: %v\tEmitter: %s\tAccount: %s\tData: %v\n", sub.EventId, sub.Emitter, sub.Account, sub.Method)
		if sub.Emitter == l.Address {
			m := relaynetwork.Message{
				Kind:    "ethbase",
				To:      sub.Account.Hex(),
				Data:    l.Data,
				EventId: sub.EventId,
			}

			return m, true
		}
	}

	return relaynetwork.Message{}, false
}

func (e *Ethbase) SubmitTx(wal *wallet.HDWallet, m relaynetwork.Message) error {
	if m.Kind != "ethbase" {
		return errors.New("tx has wrong kind")
	}

	fromAddress, err := wal.Address()
	if err != nil {
		return err
	}

	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	gasLimit := uint64(100000)

	auth, err := wal.NewKeyedTransactor()
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	subscriber := common.HexToAddress(m.To)

	tx, err := e.registry.Invoke(auth, m.EventId, subscriber, m.Data)

	log.Printf("TX sent: %s\n", tx.Hash().Hex())

	return err
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

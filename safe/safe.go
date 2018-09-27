package safe

import (
	"context"
	"encoding/json"
	"log"
	"math/big"

	common "github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	errors "github.com/pkg/errors"

	relaynetwork "github.com/planet-ethereum/relay-network"
	wallet "github.com/planet-ethereum/relay-network/wallet"
)

type SafeMessage struct {
	To             common.Address `json:"to"`
	Value          *big.Int       `json:"value"`
	Data           []byte         `json:"data"`
	Operation      uint8          `json:"operation"`
	SafeTxGas      *big.Int       `json:"safeTxGas"`
	DataGas        *big.Int       `json:"dataGas"`
	GasPrice       *big.Int       `json:"gasPrice"`
	GasToken       common.Address `json:"gasToken"`
	RefundReceiver common.Address `json:"refundReceiver"`
	SignedMessage  []byte         `json:"signedMessage"`
	Nonce          *big.Int       `json:"nonce"`
}

type SafeRelayer struct {
	client   *ethclient.Client
	contract *Safe
}

func NewSafeRelayer(client *ethclient.Client, contract *Safe) *SafeRelayer {
	r := new(SafeRelayer)
	r.client = client
	r.contract = contract
	return r
}

func (r *SafeRelayer) SubmitTx(ctx context.Context, wal *wallet.HDWallet, m relaynetwork.Message) error {
	if m.Kind != "safe" {
		return errors.New("message has wrong kind")
	}

	var sm SafeMessage
	if err := json.Unmarshal(m.Data, &sm); err != nil {
		return errors.Wrap(err, "failed to unmarshal safe message")
	}

	fromAddress, err := wal.Address()
	if err != nil {
		return err
	}

	nonce, err := r.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := r.client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	gasLimit := uint64(150000)

	auth, err := wal.NewKeyedTransactor()
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	tx, err := r.contract.ExecTransactionAndPaySubmitter(
		auth,
		sm.To,
		sm.Value,
		sm.Data,
		sm.Operation,
		sm.SafeTxGas,
		sm.DataGas,
		sm.GasPrice,
		sm.GasToken,
		sm.SignedMessage,
	)
	if err != nil {
		return errors.Wrap(err, "failed to exec transaction")
	}

	log.Printf("TX sent: %s\n", tx.Hash().Hex())

	return nil
}

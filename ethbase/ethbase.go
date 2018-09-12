package ethbase

import (
	"context"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	errors "github.com/pkg/errors"

	relaynetwork "github.com/planet-ethereum/relay-network"
)

type Subscription struct {
	EventId   [32]byte
	Emitter   common.Address
	EventName [32]byte
	Account   common.Address
	Method    [4]byte
}

func SubmitTx(client *ethclient.Client, m relaynetwork.Message) error {
	registryAddress := common.HexToAddress("0x8c42e8e89c4ee591d69fdc95803388d80290c46c")

	if m.Kind != "ethbase" {
		return errors.New("tx has wrong kind")
	}

	registry, err := NewEthbase(registryAddress, client)
	if err != nil {
		return errors.Wrap(err, "creating registry instance failed")
	}

	subscriber := common.HexToAddress(m.To)

	_, err = registry.Invoke(nil, m.EventId, subscriber, m.Data)

	return err
}

func GetSubscriptions(client *ethclient.Client, address common.Address) ([]Subscription, error) {
	/*registry, err := NewEthbase(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating registry instance failed")
	}*/

	logs, err := getPastLogs(client, address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch past logs")
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(EthbaseABI)))
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

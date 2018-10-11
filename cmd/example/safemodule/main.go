package main

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"os"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	crypto "github.com/ethereum/go-ethereum/crypto"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	errors "github.com/pkg/errors"

	//relaynetwork "github.com/planet-ethereum/relay-network"
	example "github.com/planet-ethereum/relay-network/cmd/example"
	ethbase "github.com/planet-ethereum/relay-network/ethbase"
	//ipfs "github.com/planet-ethereum/relay-network/ipfs"
	safe "github.com/planet-ethereum/relay-network/safe"
	wallet "github.com/planet-ethereum/relay-network/wallet"
)

func main() {
	// Connect to IPFS
	// ps := ipfs.NewIPFSPub("localhost:5001", "local-relay-network")

	// Connect to Ethereum client
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	//client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatalf("failed to connect to ethclient: %v", err)
	}

	moduleAddress := common.HexToAddress("0x65eB9EAE3BB8C5D6E062078Be5209D00B4a6b07f")
	moduleContract, err := ethbase.NewModule(moduleAddress, client)
	if err != nil {
		log.Fatalf("failed to instantiate safe module contract: %v\n", err)
	}
	moduleNonce, err := moduleContract.Nonce(nil)
	if err != nil {
		log.Fatalf("failed to get safe module nonce: %v\n", err)
	}
	log.Printf("Safe module nonce: %v\n", moduleNonce)

	safeMessage := createSafeMessage(client, moduleContract, moduleNonce)

	emitterAddress := common.HexToAddress("0x7b76a89ef15ef6a63e8707f841c4afd498edddc8")
	logs, err := getPastLogs(client, emitterAddress)
	if err != nil || len(logs) == 0 {
		log.Fatalf("failed to get past logs: %v\n", err)
	}

	lastLog := logs[len(logs)-1]
	logProof, err := ethbase.GenerateLogProof(context.Background(), client, lastLog)
	if err != nil {
		log.Fatalf("failed to calculate log proof: %v\n", err)
	}

	if err := submitTx(context.Background(), client, moduleContract, safeMessage, logProof, emitterAddress, lastLog.Topics[0], moduleNonce); err != nil {
		log.Fatalf("failed to submit tx: %v\n", err)
	}
}

func submitTx(
	ctx context.Context,
	client *ethclient.Client,
	contract *ethbase.Module,
	sm safe.SafeMessage,
	logProof *ethbase.LogProof,
	emitterAddress common.Address,
	eventTopic common.Hash,
	moduleNonce *big.Int,
) error {
	wal, err := wallet.NewHDWallet(os.Getenv("MNEMONIC"), "m/44'/60'/0'/0/2")
	if err != nil {
		log.Fatalf("failed to create wallet: %v\n", err)
	}

	fromAddress, err := wal.Address()
	if err != nil {
		return err
	}

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	gasLimit := uint64(350000)

	auth, err := wal.NewKeyedTransactor()
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	tx, err := contract.ExecConditional(
		auth,
		logProof.Value,
		logProof.Proof,
		logProof.Key,
		big.NewInt(int64(logProof.LogIndex)),
		logProof.Header,
		emitterAddress,
		eventTopic,
		sm.To,
		sm.Value,
		sm.Data,
		sm.Operation,
		moduleNonce,
		sm.SignedMessage,
	)
	if err != nil {
		return errors.Wrap(err, "failed to exec transaction")
	}

	log.Printf("TX sent: %s\n", tx.Hash().Hex())
	return nil
}

func createSafeMessage(client *ethclient.Client, moduleContract *ethbase.Module, moduleNonce *big.Int) safe.SafeMessage {
	// Set up wallet
	wal, err := wallet.NewHDWallet(os.Getenv("MNEMONIC"), "m/44'/60'/0'/0/2")
	if err != nil {
		log.Fatalf("failed to create wallet: %v\n", err)
	}

	fromAddress, err := wal.Address()
	if err != nil {
		log.Fatalf("failed to get account address: %v\n", err)
	}

	log.Printf("Account: %s\n", fromAddress.Hex())

	simpleStorageAddress := common.HexToAddress("0xff4b6cd1ea4011756e6729746195df097b4f697a")
	ss, err := example.NewSimpleStorage(simpleStorageAddress, client)
	if err != nil {
		log.Fatalf("failed to instantiate simple storage: %v\n", err)
	}

	res, err := ss.Get(nil)
	if err != nil {
		log.Fatalf("failed to get current value: %v\n", err)
	}
	log.Printf("Current value in storage: %v\n", res)

	parsedABI, err := abi.JSON(strings.NewReader(example.SimpleStorageABI))
	if err != nil {
		log.Fatalf("failed to parse contract ABI: %v\n", err)
	}

	value := big.NewInt(int64(5))
	data, err := parsedABI.Pack("set", value)
	if err != nil {
		log.Fatalf("failed to abi pack method and args: %v\n", err)
	}
	log.Printf("data: %s\n", hex.EncodeToString(data))

	zeroBig := new(big.Int)
	sm := safe.SafeMessage{
		To:        simpleStorageAddress,
		Value:     zeroBig,
		Data:      data,
		Operation: 0,
		Nonce:     moduleNonce,
	}

	txHash, err := moduleContract.GetTransactionHash(
		nil,
		sm.To,
		sm.Value,
		sm.Data,
		sm.Operation,
		sm.Nonce,
	)
	if err != nil {
		log.Fatalf("failed to get tx hash: %v\n", err)
	}

	log.Printf("txhash: %v\n", hex.EncodeToString(txHash[:]))
	sig, err := wal.SignHash(txHash[:])
	if err != nil {
		log.Fatalf("failed to sign tx hash: %v\n", err)
	}

	signatureNoRecoverID := sig[:len(sig)-1]
	pubKey, err := wal.PublicKey()
	if err != nil {
		log.Fatalf("failed to get public key: %v\n", err)
	}

	pubKeyBytes := crypto.FromECDSAPub(pubKey)
	ok := crypto.VerifySignature(pubKeyBytes, txHash[:], signatureNoRecoverID)
	if !ok {
		log.Fatalf("failed to verify generated signature")
	}

	sig[len(sig)-1] = sig[len(sig)-1] + 27
	sm.SignedMessage = sig

	return sm
}

func getPastLogs(client *ethclient.Client, emitterAddress common.Address) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			emitterAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

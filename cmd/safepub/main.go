package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"os"
	"strings"

	abi "github.com/ethereum/go-ethereum/accounts/abi"
	common "github.com/ethereum/go-ethereum/common"
	crypto "github.com/ethereum/go-ethereum/crypto"
	ethclient "github.com/ethereum/go-ethereum/ethclient"

	relaynetwork "github.com/planet-ethereum/relay-network"
	ipfs "github.com/planet-ethereum/relay-network/ipfs"
	safe "github.com/planet-ethereum/relay-network/safe"
	wallet "github.com/planet-ethereum/relay-network/wallet"
)

func main() {
	// Connect to IPFS
	ps := ipfs.NewIPFSPub("localhost:5001", "local-relay-network")

	// Connect to Ethereum client
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatalf("failed to connect to ethclient: %v", err)
	}

	// Set up wallet
	wal, err := wallet.NewHDWallet(os.Getenv("MNEMONIC"), "m/44'/60'/0'/0/9")
	if err != nil {
		log.Fatalf("failed to create wallet: %v\n", err)
	}

	fromAddress, err := wal.Address()
	if err != nil {
		log.Fatalf("failed to get account address: %v\n", err)
	}

	log.Printf("Account: %s\n", fromAddress.Hex())

	simpleStorageAddress := common.HexToAddress("0xff4b6cd1ea4011756e6729746195df097b4f697a")
	safeAddress := common.HexToAddress("0xd186b722d4bf23a4199ffc4c6767a5a36d136875")

	ss, err := NewSimpleStorage(simpleStorageAddress, client)
	if err != nil {
		log.Fatalf("failed to instantiate simple storage: %v\n", err)
	}

	safeContract, err := safe.NewSafe(safeAddress, client)
	if err != nil {
		log.Fatalf("failed to instantiate safe contract: %v\n", err)
	}

	safeNonce, err := safeContract.Nonce(nil)
	if err != nil {
		log.Fatalf("failed to get safe nonce: %v\n", err)
	}
	log.Printf("Safe nonce: %v\n", safeNonce)

	res, err := ss.Get(nil)
	if err != nil {
		log.Fatalf("failed to get current value: %v\n", err)
	}
	log.Printf("Current value in storage: %v\n", res)

	parsedABI, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		log.Fatalf("failed to parse contract ABI: %v\n", err)
	}

	value := big.NewInt(int64(7))
	data, err := parsedABI.Pack("set", value)
	if err != nil {
		log.Fatalf("failed to abi pack method and args: %v\n", err)
	}

	zeroBig := new(big.Int)
	zeroAddr := common.BigToAddress(zeroBig)
	txGas := big.NewInt(50000)

	sm := safe.SafeMessage{
		To:        simpleStorageAddress,
		Value:     zeroBig,
		Data:      data,
		Operation: 0,
		SafeTxGas: txGas,
		DataGas:   zeroBig,
		GasPrice:  zeroBig,
		GasToken:  zeroAddr,
		Nonce:     safeNonce,
	}

	txHash, err := safeContract.GetTransactionHash(
		nil,
		sm.To,
		sm.Value,
		sm.Data,
		sm.Operation,
		sm.SafeTxGas,
		sm.DataGas,
		sm.GasPrice,
		sm.GasToken,
		sm.Nonce,
	)
	if err != nil {
		log.Fatalf("failed to get tx hash: %v\n", err)
	}

	log.Printf("txhash: %v\n", txHash)
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

	rawSafeMsg, err := json.Marshal(sm)
	if err != nil {
		log.Fatalf("failed to marshalize safe message: %v\n", err)
	}

	m := relaynetwork.Message{
		Kind: "safe",
		To:   safeAddress.Hex(),
		Data: rawSafeMsg,
	}

	msg, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("failed to marshalize message: %v\n", err)
	}

	if err = ps.Publish(string(msg)); err != nil {
		log.Fatalf("failed to publish to pubsub: %v\n", err)
	}
}

package ethbase

import (
	"bytes"
	"context"
	"log"

	types "github.com/ethereum/go-ethereum/core/types"
	crypto "github.com/ethereum/go-ethereum/crypto"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	ethdb "github.com/ethereum/go-ethereum/ethdb"
	rlp "github.com/ethereum/go-ethereum/rlp"
	trie "github.com/ethereum/go-ethereum/trie"
	errors "github.com/pkg/errors"
)

// LogProof contains everything that's necessary
// to verify the inclusion of an event in a block
// on chain.
type LogProof struct {
	Header   *types.Header
	Key      []byte
	Value    []byte
	LogIndex uint
	Proof    trie.DatabaseReader
}

// GenerateLogProof generates a LogProof instance
// for a given log.
func GenerateLogProof(ctx context.Context, client *ethclient.Client, l types.Log) (*LogProof, error) {
	block, err := client.BlockByHash(ctx, l.BlockHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get blockByHash")
	}

	header := block.Header()
	receipts, err := getBlockReceipts(ctx, client, block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block receipts")
	}

	tr, err := trieFromReceipts(receipts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create trie out of receipts")
	}

	// Generate MTP.
	proof := ethdb.NewMemDatabase()
	key, err := rlp.EncodeToBytes(l.TxIndex)
	if err != nil {
		return nil, err
	}

	if it := trie.NewIterator(tr.NodeIterator(key)); it.Next() && bytes.Equal(key, it.Key) {
		for _, p := range it.Prove() {
			proof.Put(crypto.Keccak256(p), p)
		}
	}

	logRawBuf := new(bytes.Buffer)
	if err := l.EncodeRLP(logRawBuf); err != nil {
		return nil, err
	}

	logProof := &LogProof{
		Header:   header,
		Key:      key,
		Proof:    proof,
		LogIndex: uint(0), // TODO: Fix log index (l.Index returns weird number)
		Value:    logRawBuf.Bytes(),
	}

	return logProof, nil
}

// VerifyLogProof simulates verification of log proofs on-chain.
// It only assumes blockNum -> blockHash is accessible.
func VerifyLogProof(ctx context.Context, client *ethclient.Client, p *LogProof) (bool, error) {
	block, err := client.BlockByNumber(ctx, p.Header.Number)
	if err != nil {
		return false, err
	}

	if block.Hash() != p.Header.Hash() {
		log.Printf("Invalid Proof: Proof Header hash: %s\tBlock Hash: %s\n", block.Hash().Hex(), p.Header.Hash().Hex())
		return false, nil
	}

	// Recover rlp encoded receipt from proof and receipt
	// trie root.
	receiptRaw, _, err := trie.VerifyProof(p.Header.ReceiptHash, p.Key, p.Proof)
	lRaw, err := logFromReceipt(receiptRaw, p.LogIndex)
	if err != nil {
		return false, err
	}

	if !bytes.Equal(lRaw, p.Value) {
		log.Printf("invalid proof: computed leaf doesn't match\n")
		return false, nil
	}

	return true, nil
}

func getBlockReceipts(ctx context.Context, client *ethclient.Client, block *types.Block) ([]*types.Receipt, error) {
	txes := block.Transactions()
	receipts := make([]*types.Receipt, len(txes))

	for i, tx := range txes {
		r, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, errors.Wrap(err, "failed to get tx receipt")
		}

		receipts[i] = r
	}

	return receipts, nil
}

func trieFromReceipts(receipts []*types.Receipt) (*trie.Trie, error) {
	tr := new(trie.Trie)

	for i, r := range receipts {
		path, err := rlp.EncodeToBytes(uint(i))
		if err != nil {
			return nil, errors.Wrap(err, "failed to encode index")
		}

		rawReceipt, err := encodeReceipt(r)
		if err != nil {
			return nil, errors.Wrap(err, "failed to encode receipt")
		}

		tr.Update(path, rawReceipt)
	}

	return tr, nil
}

func logFromReceipt(r []byte, i uint) ([]byte, error) {
	var receipt types.Receipt
	if err := rlp.DecodeBytes(r, &receipt); err != nil {
		return nil, err
	}

	// Extract log from receipt.
	l := receipt.Logs[i]
	lRawBuf := new(bytes.Buffer)
	if err := l.EncodeRLP(lRawBuf); err != nil {
		return nil, err
	}

	return lRawBuf.Bytes(), nil
}

func encodeReceipt(r *types.Receipt) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := r.EncodeRLP(buf); err != nil {
		return nil, errors.Wrap(err, "failed to encode receipt")
	}

	return buf.Bytes(), nil
}

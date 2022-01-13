package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Transactions          []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height int
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, preBlockHash []byte, height int) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:          transactions,
		PrevBlockHash: preBlockHash,
		Height: height,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	_ = encoder.Encode(b)
	return result.Bytes()

}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	_ = decoder.Decode(&block)
	return &block
}


// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	tree := NewMerkleTree(transactions)
	return tree.RootNode.Data
}


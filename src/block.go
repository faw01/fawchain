package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// block represents the basic data unit in the blockchain
type Block struct {
	Timestamp     string
	Data          float64 // represents the base price
	PrevBlockHash string
	Hash          string
	Seed          string
	Price         float64 // calculated price based on the algorithm
}

// sethash computes the hash of the block based on its contents
func (b *Block) SetHash() {
	timestamp := []byte(b.Timestamp)
	dataString := fmt.Sprintf("%f", b.Data)
	headers := bytes.Join([][]byte{[]byte(b.PrevBlockHash), []byte(dataString), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hex.EncodeToString(hash[:])
}

// creates a new block in the blockchain
func NewBlock(data float64, prevBlockHash []byte, seed string) *Block {
	block := &Block{
		Timestamp:     time.Now().Format("2006-01-02 15:04:05"),
		Data:          data,
		PrevBlockHash: string(prevBlockHash),
		Seed:          seed,
	}
	block.Price = determinePrice(seed, data) // calculate price based on seed and base price
	block.SetHash()
	return block
}

// creates the first block in the blockchain
func NewGenesisBlock() *Block {
	return &Block{
		Timestamp:     time.Now().Format("2006-01-02 15:04:05"),
		Data:          0,
		PrevBlockHash: "Genesis",
		Hash:          "1",
		Seed:          "1",
		Price:         0,
	}
}

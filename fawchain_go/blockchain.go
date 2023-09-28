package main

import (
	"crypto/rand"
	"encoding/hex"
)

// the blockchain
type Blockchain struct {
	blocks []*Block
}

// adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data float64) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	seed := generateRandomSeed() // generate a random seed for price calculation
	newBlock := NewBlock(data, []byte(prevBlock.Hash), seed)
	bc.blocks = append(bc.blocks, newBlock)
}

// creates a new blockchain with a genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := NewGenesisBlock()
	return &Blockchain{[]*Block{genesisBlock}}
}

// produces a random seed for price calculation
func generateRandomSeed() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}

package main

import (
	"fmt"
	"strings"
)

type Blockchain struct {
	TransactionPool []string
	Chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce uint64, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.Chain = append(bc.Chain, b)
	return b
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	lastBlock := &Block{}
	for i := uint64(0); i < 3; i++ {
		lastBlock = bc.CreateBlock(i, lastBlock.Hash())
	}
	return bc
}

func (bc *Blockchain) Print() {
	pattern := strings.Repeat("=", 25)
	for i, block := range bc.Chain {
		fmt.Printf("%s Block %d %s\n", pattern, i, pattern)
		block.Print()
	}
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

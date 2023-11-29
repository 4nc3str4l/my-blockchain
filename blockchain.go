package main

import (
	"fmt"
	"strings"
)

type Blockchain struct {
	TransactionPool []string
	Chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce uint64, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.Chain = append(bc.Chain, b)
	return b
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	for i := uint64(0); i < 3; i++ {
		bc.CreateBlock(i, fmt.Sprintf("Hash %d", i))
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

func InitBlockchain() {
	bc := NewBlockchain()
	bc.Print()
}

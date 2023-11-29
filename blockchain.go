package main

import (
	"fmt"
)

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce uint64, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
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
	for i, block := range bc.chain {
		fmt.Printf("Block %d \n", i)
		block.Print()
	}
}

func InitBlockchain() {
	bc := NewBlockchain()
	bc.Print()
}
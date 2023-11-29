package main

import (
	"fmt"
	"math/big"
	"strings"
)

type Blockchain struct {
	TransactionPool []*Transaction
	Chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce uint64, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, b)
	bc.TransactionPool = []*Transaction{}
	return b
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)

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

func (bc *Blockchain) AddTransaction(sender string, recipient string, value *big.Int) {
	t := NewTransaction(sender, recipient, value)
	bc.TransactionPool = append(bc.TransactionPool, t)
}

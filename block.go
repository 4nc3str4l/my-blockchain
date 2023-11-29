package main

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        uint64
	prevHash     string
	timestamp    uint64
	transactions []string
}

func NewBlock(nonce uint64, previousHash string) *Block {
	b := new(Block)
	b.timestamp = uint64(time.Now().UnixNano())
	b.nonce = nonce
	b.prevHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp \t%d\n", b.timestamp)
	fmt.Printf("nonce \t\t%d\n", b.nonce)
	fmt.Printf("previous_hash \t%s\n", b.prevHash)
	fmt.Printf("transactions \t%s\n", b.transactions)
}

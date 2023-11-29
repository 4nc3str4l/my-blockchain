package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        uint64
	prevHash     [32]byte
	timestamp    uint64
	transactions []*Transaction
}

func NewBlock(nonce uint64, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		nonce:        nonce,
		prevHash:     previousHash,
		timestamp:    uint64(time.Now().UnixNano()),
		transactions: transactions,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(&struct {
		Nonce        uint64         `json:"nonce"`
		PrevHash     [32]byte       `json:"previous_hash"`
		Timestamp    uint64         `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
	return sha256.Sum256([]byte(m))
}

func (b *Block) Print() {
	fmt.Printf("timestamp \t%d\n", b.timestamp)
	fmt.Printf("nonce \t\t%d\n", b.nonce)
	fmt.Printf("previous_hash \t%x\n", b.prevHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

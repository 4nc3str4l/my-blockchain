package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Nonce        uint64   `json:"nonce"`
	PrevHash     [32]byte `json:"previous_hash"`
	Timestamp    uint64   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

func NewBlock(nonce uint64, previousHash [32]byte) *Block {
	b := new(Block)
	b.Timestamp = uint64(time.Now().UnixNano())
	b.Nonce = nonce
	b.PrevHash = previousHash
	return b
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) Print() {
	fmt.Printf("timestamp \t%d\n", b.Timestamp)
	fmt.Printf("nonce \t\t%d\n", b.Nonce)
	fmt.Printf("previous_hash \t%x\n", b.PrevHash)
	fmt.Printf("transactions \t%s\n", b.Transactions)
	b.Hash()
}

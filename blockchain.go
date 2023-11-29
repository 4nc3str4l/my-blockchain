package main

import (
	"fmt"
	"math/big"
	"strings"
)

const MINING_DIFFICULTY = 5

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
	b := &Block{}
	bc.CreateBlock(0, b.Hash())
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

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.TransactionPool {
		transactions = append(transactions,
			NewTransaction(t.senderAddr, t.recipientAddr, t.value))
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce uint64, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() uint64 {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().prevHash
	nonce := uint64(0)
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

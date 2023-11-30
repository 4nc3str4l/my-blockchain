package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
	COINBASE_ADDRESS  = "CoinbaseTransaction"
	MINING_REWARDS    = "5000000000" // sats (50 btc)
)

type Blockchain struct {
	TransactionPool []*Transaction
	Chain           []*Block
	Address         string
}

func (bc *Blockchain) CreateBlock(nonce uint64, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, b)
	bc.TransactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, s *Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func NewBlockchain(addr string) *Blockchain {
	bc := new(Blockchain)
	bc.Address = addr
	b := &Block{}
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) Print() {
	pattern := strings.Repeat("=", 25)
	for i, block := range bc.Chain {
		log.Printf("%s Block %d %s\n", pattern, i, pattern)
		block.Print()
	}
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value *big.Int, senderPublicKey *ecdsa.PublicKey, s *Signature) bool {
	t := NewTransaction(sender, recipient, value)

	// If it is a mining reward we don't need to verify the signature
	if sender == COINBASE_ADDRESS {
		bc.TransactionPool = append(bc.TransactionPool, t)
		return true
	}

	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		if bc.ComputeBalance(sender).Cmp(value) < 0 {
			log.Println("ERROR: Not enough balance in a wallet")
			return false
		}
		bc.TransactionPool = append(bc.TransactionPool, t)
		return true
	} else {
		log.Println("ERROR: Verify Transaction")
	}
	return false

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

func (bc *Blockchain) Mine() bool {
	// Here we include our own reward
	reward, _ := new(big.Int).SetString(MINING_REWARDS, 10)
	bc.AddTransaction(COINBASE_ADDRESS, bc.Address, reward, nil, nil)

	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=⛏️, status=success🔵")
	return true
}

func (bc *Blockchain) ComputeBalance(blockchainAddress string) *big.Int {
	totalAmmount := new(big.Int)
	totalAmmount, _ = totalAmmount.SetString("0", 10)
	for _, b := range bc.Chain {
		for _, t := range b.transactions {
			value := t.value
			if blockchainAddress == t.recipientAddr {
				totalAmmount.Add(totalAmmount, value)
			}

			if blockchainAddress == t.senderAddr {
				totalAmmount.Sub(totalAmmount, value)
			}
		}
	}
	return totalAmmount
}

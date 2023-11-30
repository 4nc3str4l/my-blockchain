package main

import (
	"log"
	"math/big"
)

func main() {
	quantity, success := new(big.Int).SetString("100000000", 10)
	if !success {
		panic("Could not parse quantity")
	}

	walletMiner := NewWallet()
	walletA := NewWallet()
	walletB := NewWallet()

	t := NewUnsignedTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), quantity)

	// Blockchain
	blockchain := NewBlockchain(walletMiner.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), quantity, walletA.PublicKey(), t.GenerateSignature())
	log.Println("Added: ", isAdded)
}

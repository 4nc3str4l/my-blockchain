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

	blockchain.Mine()
	blockchain.Print()

	log.Printf("A %d\n", blockchain.ComputeBalance(walletA.BlockchainAddress()))
	log.Printf("B %d\n", blockchain.ComputeBalance(walletB.BlockchainAddress()))
	log.Printf("M %d\n", blockchain.ComputeBalance(walletMiner.BlockchainAddress()))
}

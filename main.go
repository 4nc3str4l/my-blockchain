package main

import (
	"log"
	"math/big"
)

func main() {
	quantity, success := new(big.Int).SetString("100000000", 10)
	w := NewWallet()
	log.Println(w.PrivateKeyStr())
	log.Println(w.PublicKeyStr())
	log.Println(w.BlockchainAddress())

	t := NewUnsignedTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", quantity)
	log.Printf("signature %s \n", t.GenerateSignature())

	bc := NewBlockchain("My Address")
	for i := 0; i < 3; i++ {

		if !success {
			panic("Could not parse quantity")
		}
		bc.AddTransaction("A", "B", quantity)
		bc.Mine()
	}
	bc.Print()

	// For now I allow negative balances (I'll tacke this when creating wallets)
	log.Println(bc.ComputeBalance("A"))
	log.Println(bc.ComputeBalance("B"))
	log.Println(bc.ComputeBalance("My Address"))

}

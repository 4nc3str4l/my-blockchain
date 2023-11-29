package main

import (
	"math/big"
)

func main() {
	bc := NewBlockchain()
	lastBlock := bc.LastBlock()
	var total_supply = new(big.Int)
	total_supply.SetString("21000000000000000000", 10)
	bc.AddTransaction("0x00", "0x00", total_supply)
	for i := 0; i < 3; i++ {
		bc.AddTransaction("0x00", "0x00", total_supply)
		nonce := bc.ProofOfWork()
		lastBlock = bc.CreateBlock(nonce, lastBlock.Hash())
	}
	bc.Print()
}

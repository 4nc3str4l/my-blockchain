package main

import "math/big"

func main() {
	bc := NewBlockchain()
	lastBlock := &Block{}
	var total_supply = new(big.Int)
	total_supply.SetString("21000000000000000000", 10)
	bc.AddTransaction("0x00", "0x00", total_supply)
	for i := uint64(0); i < 3; i++ {
		lastBlock = bc.CreateBlock(i, lastBlock.Hash())
	}
	bc.Print()
}

package main

import "math/big"

func main() {
	bc := NewBlockchain("My Address")
	for i := 0; i < 3; i++ {
		quantity, success := new(big.Int).SetString("100000000", 10)
		if !success {
			panic("Could not parse quantity")
		}
		bc.AddTransaction("0x00", "0x00", quantity)
		bc.Mining()
	}
	bc.Print()
}

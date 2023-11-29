package main

import (
	"log"
	"math/big"
)

func main() {
	bc := NewBlockchain("My Address")
	for i := 0; i < 3; i++ {
		quantity, success := new(big.Int).SetString("100000000", 10)
		if !success {
			panic("Could not parse quantity")
		}
		bc.AddTransaction("A", "B", quantity)
		bc.Mine()
	}
	bc.Print()
	log.Println(bc.ComputeBalance("A"))
}

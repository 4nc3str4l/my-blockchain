package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type Transaction struct {
	senderAddr    string
	recipientAddr string
	value         *big.Int
}

func NewTransaction(sender string, recipient string, value *big.Int) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_address    %s\n", t.senderAddr)
	fmt.Printf(" recipient_address %s\n", t.recipientAddr)
	fmt.Printf(" value             %d\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAddr    string   `json:"sender_address"`
		RecipientAddr string   `json:"recipient_address"`
		Value         *big.Int `json:"value"`
	}{
		SenderAddr:    t.senderAddr,
		RecipientAddr: t.recipientAddr,
		Value:         t.value,
	})
}

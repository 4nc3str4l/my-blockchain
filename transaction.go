package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	SenderAddr    string `json:"sender_address"`
	RecipientAddr string `json:"recipient_address"`
	Value         uint64 `json:"value"`
}

func NewTransaction(sender string, recipient string, value uint64) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_address    %s\n", t.SenderAddr)
	fmt.Printf(" recipient_address %s\n", t.RecipientAddr)
	fmt.Printf(" value             %d\n", t.Value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

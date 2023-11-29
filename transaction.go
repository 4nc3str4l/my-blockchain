package main

import (
	"encoding/json"
	"log"
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
	log.Printf("%s\n", strings.Repeat("-", 40))
	log.Printf(" sender_address    %s\n", t.senderAddr)
	log.Printf(" recipient_address %s\n", t.recipientAddr)
	log.Printf(" value             %d\n", t.value)
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

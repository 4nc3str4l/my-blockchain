package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
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

/// NEw transaction logic

type UnsignedTransaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      *big.Int
}

func NewUnsignedTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey,
	sender string, recipient string, value *big.Int) *UnsignedTransaction {
	return &UnsignedTransaction{privateKey, publicKey, sender, recipient, value}
}

func (t *UnsignedTransaction) GenerateSignature() *Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &Signature{r, s}
}

func (t *UnsignedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string   `json:"sender_blockchain_address"`
		Recipient string   `json:"recipient_blockchain_adddress"`
		Value     *big.Int `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

// represents a transfer of value between two parties.
type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
	Signature []byte // cryptographic signature to validate the transaction
}

// creates and signs a new transaction.
func CreateTransaction(sender string, recipient string, amount int, privateKey *rsa.PrivateKey) *Transaction {
	t := &Transaction{Sender: sender, Recipient: recipient, Amount: amount}
	t.signTransaction(privateKey) // sign the transaction
	return t
}

// signs the transaction with the sender's private key.
func (t *Transaction) signTransaction(privateKey *rsa.PrivateKey) {
	hashed := sha256.Sum256([]byte(t.Sender + t.Recipient + string(t.Amount)))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		log.Fatalf("Error signing the transaction: %s", err.Error())
	}
	t.Signature = signature
}

// verifies the transaction's signature.
func (t *Transaction) isTransactionValid(publicKey *rsa.PublicKey) bool {
	hashed := sha256.Sum256([]byte(t.Sender + t.Recipient + string(t.Amount)))
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], t.Signature)
	return err == nil
}

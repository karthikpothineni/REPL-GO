package command

import (
	"REPL-GO/internal/datastore"
)

// currentTransaction represents current running transaction
var currentTransaction *Transaction

// Transaction is a struct representing REPL transactions
type Transaction struct {
	data            map[string]string
	prevTransaction *Transaction
}

// NewTransaction returns the instance of Transaction
func NewTransaction() *Transaction {
	return &Transaction{
		data:            nil,
		prevTransaction: nil,
	}
}

// GetCurrentTransaction returns the current running transaction
func GetCurrentTransaction() *Transaction {
	return currentTransaction
}

// Start updates the transaction data and previous transaction link
func (transaction *Transaction) Start(store *datastore.Datastore) {
	// get current running transaction
	currTransaction := GetCurrentTransaction()

	// update new transaction
	if currTransaction != nil {
		transaction.data = copyMap(currTransaction.data)
		transaction.prevTransaction = currTransaction
	} else {
		transaction.data = copyMap(store.Data)
	}

	currentTransaction = transaction
}

// Commit assigns current transaction data to previous transaction or datastore
func (transaction *Transaction) Commit(store *datastore.Datastore) {
	currentData := transaction.data

	if transaction.prevTransaction != nil {
		transaction.prevTransaction.data = copyMap(currentData)
	} else {
		store.Data = copyMap(currentData)
	}

	transaction.updateCurrentTransaction()
}

// Abort aborts the transaction changes and updates current transaction
func (transaction *Transaction) Abort() {
	transaction.updateCurrentTransaction()
}

// updateCurrentTransaction updates the current transaction
func (transaction *Transaction) updateCurrentTransaction() {
	if transaction.prevTransaction != nil {
		currentTransaction = transaction.prevTransaction
	} else {
		currentTransaction = nil
	}
}

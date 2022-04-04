package command

import (
	"fmt"

	"REPL-GO/internal/datastore"
)

var currentTransaction *Transaction

type Transaction struct {
	data            map[string]string
	prevTransaction *Transaction
}

func NewTransaction(data map[string]string, prevTransaction *Transaction) *Transaction {
	return &Transaction{
		data:            data,
		prevTransaction: prevTransaction,
	}
}

func StartTransaction(store *datastore.Datastore) {
	var transaction *Transaction

	// create new transaction
	currTransaction := GetCurrentTransaction()
	if currTransaction != nil {
		transaction = NewTransaction(copyMap(currTransaction.data), currTransaction)
	} else {
		transaction = NewTransaction(copyMap(store.Data), nil)
	}

	currentTransaction = transaction
}

func CommitTransaction(transaction *Transaction, store *datastore.Datastore) {
	if transaction == nil {
		fmt.Println("there is no transaction to commit")
		return
	}

	currentData := transaction.data

	// assign current transaction data to previous transaction or datastore
	if transaction.prevTransaction != nil {
		transaction.prevTransaction.data = copyMap(currentData)
	} else {
		store.Data = copyMap(currentData)
	}

	UpdateCurrentTransaction(transaction)
}

func AbortTransaction(transaction *Transaction) {
	if transaction == nil {
		fmt.Println("there is no transaction to abort")
		return
	}

	UpdateCurrentTransaction(transaction)
}

func UpdateCurrentTransaction(transaction *Transaction) {
	if transaction.prevTransaction != nil {
		currentTransaction = transaction.prevTransaction
	} else {
		currentTransaction = nil
	}
}

func GetCurrentTransaction() *Transaction {
	return currentTransaction
}

package command

import (
	"fmt"

	"REPL-GO/internal/datastore"
)

func WriteData(transaction *Transaction, store *datastore.Datastore, key, value string) {
	if transaction == nil {
		store.Data[key] = value
	} else {
		transaction.data[key] = value
	}
}

func ReadAndPrintData(transaction *Transaction, store *datastore.Datastore, key string) {
	var value string

	var ok bool

	if transaction == nil {
		value, ok = store.Data[key]
	} else {
		value, ok = transaction.data[key]
	}

	if !ok {
		fmt.Printf("key not found: %s \n", key)
		return
	}

	fmt.Println(value)
}

func DeleteData(transaction *Transaction, store *datastore.Datastore, key string) {
	if transaction == nil {
		delete(store.Data, key)
	} else {
		delete(transaction.data, key)
	}
}

package command

import (
	"fmt"
)

func WriteData(transaction *Transaction, datastore *Datastore, key, value string) {
	if transaction == nil {
		datastore.data[key] = value
	} else {
		transaction.data[key] = value
	}
}

func ReadAndPrintData(transaction *Transaction, datastore *Datastore, key string) {
	var value string
	var ok bool

	if transaction == nil {
		value, ok = datastore.data[key]
	} else {
		value, ok = transaction.data[key]
	}

	if !ok {
		fmt.Printf("key not found: %s \n", key)
		return
	}

	fmt.Println(value)
}

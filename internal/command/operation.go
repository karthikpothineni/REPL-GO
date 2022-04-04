package command

import (
	"REPL-GO/internal/datastore"
)

// WriteData adds/updates the value for a key in transaction or store data
func WriteData(transaction *Transaction, store *datastore.Datastore, key, value string) {
	if transaction == nil {
		store.Data[key] = value
	} else {
		transaction.data[key] = value
	}
}

// ReadData reads the value for a key from transaction or store data
func ReadData(transaction *Transaction, store *datastore.Datastore, key string) (string, bool) {
	var value string

	var ok bool

	if transaction == nil {
		value, ok = store.Data[key]
	} else {
		value, ok = transaction.data[key]
	}

	return value, ok
}

// DeleteData deletes a key entry from transaction or store data
func DeleteData(transaction *Transaction, store *datastore.Datastore, key string) {
	if transaction == nil {
		delete(store.Data, key)
	} else {
		delete(transaction.data, key)
	}
}

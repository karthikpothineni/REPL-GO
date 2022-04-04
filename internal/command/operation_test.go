package command

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"REPL-GO/internal/datastore"
)

// TestWriteDataNoTransaction tests write data when there is no transaction
func TestWriteDataNoTransaction(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	expectedData := map[string]string{
		"a": "hello",
	}

	WriteData(nil, store, "a", "hello")
	check.EqualValues(expectedData, store.Data)
}

// TestWriteDataWithTransaction tests write data when there is a transaction
func TestWriteDataWithTransaction(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	expectedData := map[string]string{
		"a": "hello",
	}
	transaction := NewTransaction()

	transaction.Start(store)

	WriteData(transaction, store, "a", "hello")
	check.EqualValues(expectedData, transaction.data)
}

// TestReadDataNoTransactionKeyFound tests read data when there is no transaction and key is present
func TestReadDataNoTransactionKeyFound(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	WriteData(nil, store, "a", "hello")

	value, ok := ReadData(nil, store, "a")
	check.Equal("hello", value)
	check.Equal(true, ok)
}

// TestReadDataNoTransactionKeyNotFound tests read data when there is no transaction and key is not present
func TestReadDataNoTransactionKeyNotFound(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	WriteData(nil, store, "b", "hello")

	value, ok := ReadData(nil, store, "a")
	check.Equal("", value)
	check.Equal(false, ok)
}

// TestReadDataWithTransactionKeyFound tests read data when there is a transaction and key is present
func TestReadDataWithTransactionKeyFound(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	transaction := NewTransaction()

	transaction.Start(store)
	WriteData(transaction, store, "a", "hello")

	value, ok := ReadData(transaction, store, "a")
	check.Equal("hello", value)
	check.Equal(true, ok)
}

// TestReadDataWithTransactionKeyNotFound tests read data when there is a transaction and key is not present
func TestReadDataWithTransactionKeyNotFound(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	transaction := NewTransaction()

	transaction.Start(store)

	value, ok := ReadData(transaction, store, "c")
	check.Equal("", value)
	check.Equal(false, ok)
}

// TestDeleteDataNoTransaction tests delete data when there is no transaction
func TestDeleteDataNoTransaction(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	expectedData := map[string]string{}

	DeleteData(nil, store, "a")
	check.EqualValues(expectedData, store.Data)
}

// TestDeleteDataWithTransaction tests delete data when there is a transaction
func TestDeleteDataWithTransaction(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	expectedData := map[string]string{}
	transaction := NewTransaction()

	transaction.Start(store)

	DeleteData(transaction, store, "a")
	check.EqualValues(expectedData, transaction.data)
}

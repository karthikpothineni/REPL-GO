package command

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"REPL-GO/internal/datastore"
)

// TestNewTransaction tests the creation of new transaction
func TestNewTransaction(t *testing.T) {
	check := assert.New(t)

	transaction := NewTransaction()
	check.Nil(transaction.prevTransaction)
	check.Nil(transaction.data)
}

// TestTransaction_StartNonNestedTransaction tests transaction start when there is a non nested transaction
func TestTransaction_StartNonNestedTransaction(t *testing.T) {
	check := assert.New(t)
	currentTransaction = nil
	expectedData := map[string]string{
		"a": "hello",
	}

	store := datastore.NewDatastore()
	WriteData(nil, store, "a", "hello")

	transaction := NewTransaction()
	transaction.Start(store)

	check.EqualValues(expectedData, transaction.data)
	check.Nil(transaction.prevTransaction)
}

// TestTransaction_StartNestedTransaction tests transaction start when there is a nested transaction
func TestTransaction_StartNestedTransaction(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()
	expectedData := map[string]string{
		"a": "hello",
	}
	transaction := NewTransaction()
	transaction.Start(store)
	WriteData(transaction, store, "a", "hello")

	nestedTransaction := NewTransaction()
	nestedTransaction.Start(store)
	check.EqualValues(expectedData, nestedTransaction.data)
	check.NotNil(nestedTransaction.prevTransaction)
}

// TestTransaction_CommitNonNestedTransaction tests transaction commit when there is a non nested transaction
func TestTransaction_CommitNonNestedTransaction(t *testing.T) {
	check := assert.New(t)
	currentTransaction = nil
	expectedData := map[string]string{
		"b": "hi",
	}

	store := datastore.NewDatastore()
	WriteData(nil, store, "b", "hello")

	transaction := NewTransaction()
	transaction.Start(store)
	WriteData(transaction, store, "b", "hi")
	transaction.Commit(store)

	check.EqualValues(expectedData, store.Data)
	check.Nil(GetCurrentTransaction())
}

// TestTransaction_CommitNestedTransaction tests transaction commit when there is a nested transaction
func TestTransaction_CommitNestedTransaction(t *testing.T) {
	check := assert.New(t)
	expectedData := map[string]string{
		"a": "hi",
		"b": "hi",
	}

	store := datastore.NewDatastore()
	WriteData(nil, store, "a", "hello")

	transaction := NewTransaction()
	transaction.Start(store)
	WriteData(transaction, store, "a", "hi")

	nestedTransaction := NewTransaction()
	nestedTransaction.Start(store)
	WriteData(nestedTransaction, store, "b", "hi")
	nestedTransaction.Commit(store)

	check.EqualValues(expectedData, transaction.data)
	check.NotNil(GetCurrentTransaction())
}

// TestTransaction_AbortNonNestedTransaction tests transaction abort when there is a non nested transaction
func TestTransaction_AbortNonNestedTransaction(t *testing.T) {
	check := assert.New(t)
	currentTransaction = nil
	expectedData := map[string]string{
		"a": "hello",
	}

	store := datastore.NewDatastore()
	WriteData(nil, store, "a", "hello")

	transaction := NewTransaction()
	transaction.Start(store)
	WriteData(transaction, store, "a", "hi")
	transaction.Abort()

	check.EqualValues(expectedData, store.Data)
	check.Nil(GetCurrentTransaction())
}

// TestTransaction_AbortNestedTransaction tests transaction abort when there is a nested transaction
func TestTransaction_AbortNestedTransaction(t *testing.T) {
	check := assert.New(t)
	currentTransaction = nil
	expectedData := map[string]string{
		"a": "hi",
	}

	store := datastore.NewDatastore()
	WriteData(nil, store, "a", "hello")

	transaction := NewTransaction()
	transaction.Start(store)
	WriteData(transaction, store, "a", "hi")

	nestedTransaction := NewTransaction()
	nestedTransaction.Start(store)
	WriteData(nestedTransaction, store, "b", "hi")
	nestedTransaction.Abort()

	check.EqualValues(expectedData, transaction.data)
	check.NotNil(GetCurrentTransaction())
}

// TestTransaction_updateCurrentNonNestedTransaction tests current transaction update when there is a non nested transaction
func TestTransaction_updateCurrentNonNestedTransaction(t *testing.T) {
	check := assert.New(t)
	currentTransaction = nil
	store := datastore.NewDatastore()

	transaction := NewTransaction()
	transaction.Start(store)
	transaction.updateCurrentTransaction()

	check.Nil(GetCurrentTransaction())
}

// TestTransaction_updateCurrentNestedTransaction tests current transaction update when there is a nested transaction
func TestTransaction_updateCurrentNestedTransaction(t *testing.T) {
	check := assert.New(t)
	store := datastore.NewDatastore()

	transaction := NewTransaction()
	transaction.Start(store)

	nestedTransaction := NewTransaction()
	nestedTransaction.Start(store)
	nestedTransaction.updateCurrentTransaction()

	check.NotNil(GetCurrentTransaction())
}

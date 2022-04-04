package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewDataStore tests the creation of new datastore
func TestNewDataStore(t *testing.T) {
	check := assert.New(t)
	expectedData := make(map[string]string)

	store := NewDatastore()
	check.Equal(expectedData, store.Data)
}

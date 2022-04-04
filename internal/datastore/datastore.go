package datastore

// Datastore is a struct representing the in-memory key/value storage
type Datastore struct {
	Data map[string]string
}

// NewDatastore returns the instance of Datastore
func NewDatastore() *Datastore {
	return &Datastore{
		Data: make(map[string]string),
	}
}

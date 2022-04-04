package datastore

type Datastore struct {
	Data map[string]string
}

func NewDataStore() *Datastore {
	return &Datastore{
		Data: make(map[string]string),
	}
}

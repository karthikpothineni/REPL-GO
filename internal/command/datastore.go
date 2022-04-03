package command

type Datastore struct {
	data map[string]string
}

func NewDataStore() *Datastore {
	return &Datastore{
		data: make(map[string]string),
	}
}

func (d *Datastore) GetData() map[string]string {
	return d.data
}

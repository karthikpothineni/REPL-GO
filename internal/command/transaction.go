package command

var lastTransaction *Transaction

type Transaction struct {
	data map[string]string
}

func NewTransaction(data map[string]string) *Transaction {
	return &Transaction{
		data: data, // copy data from other main data to here
	}
}

func StartTransaction(datastore *Datastore) *Transaction {
	var transaction *Transaction

	prevTransaction := GetLastTransaction()
	if prevTransaction != nil {
		transaction = NewTransaction(copyMap(prevTransaction.data))
	} else {
		transaction = NewTransaction(copyMap(datastore.data))
	}

	lastTransaction = transaction

	return transaction
}

func GetLastTransaction() *Transaction {
	return lastTransaction
}

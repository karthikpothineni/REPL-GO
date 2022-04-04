package command

import (
	"fmt"
	"os"

	"REPL-GO/internal/datastore"
)

// command list
const (
	Write  = "write"
	Read   = "read"
	Delete = "delete"
	Start  = "start"
	Commit = "commit"
	Abort  = "abort"
	Quit   = "quit"
)

// Execute executes the operations for commands
func Execute(store *datastore.Datastore, input []string) {
	currTransaction := GetCurrentTransaction()

	switch input[0] {
	case Write:
		WriteData(currTransaction, store, input[1], input[2])
	case Read:
		value, ok := ReadData(currTransaction, store, input[1])
		if !ok {
			fmt.Printf("key not found: %s \n", input[1])
			return
		}

		fmt.Println(value)
	case Delete:
		DeleteData(currTransaction, store, input[1])
	case Start:
		newTransaction := NewTransaction()
		newTransaction.Start(store)
	case Commit:
		if currTransaction == nil {
			fmt.Println("there is no transaction to commit")
			return
		}

		currTransaction.Commit(store)
	case Abort:
		if currTransaction == nil {
			fmt.Println("there is no transaction to abort")
			return
		}

		currTransaction.Abort()
	case Quit:
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}

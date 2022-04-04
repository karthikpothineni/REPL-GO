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

func Execute(store *datastore.Datastore, input []string) {
	switch input[0] {
	case Write:
		WriteData(GetCurrentTransaction(), store, input[1], input[2])
	case Read:
		ReadAndPrintData(GetCurrentTransaction(), store, input[1])
	case Delete:
		DeleteData(GetCurrentTransaction(), store, input[1])
	case Start:
		StartTransaction(store)
	case Commit:
		CommitTransaction(GetCurrentTransaction(), store)
	case Abort:
		AbortTransaction(GetCurrentTransaction())
	case Quit:
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}

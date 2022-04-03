package command

import (
	"fmt"
	"os"
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

func Execute(datastore *Datastore, input []string) {
	switch input[0] {
	case Write:
		WriteData(GetCurrentTransaction(), datastore, input[1], input[2])
	case Read:
		ReadAndPrintData(GetCurrentTransaction(), datastore, input[1])
	case Delete:
		DeleteData(GetCurrentTransaction(), datastore, input[1])
	case Start:
		StartTransaction(datastore)
	case Commit:
		CommitTransaction(GetCurrentTransaction(), datastore)
	case Abort:
		AbortTransaction(GetCurrentTransaction())
	case Quit:
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}

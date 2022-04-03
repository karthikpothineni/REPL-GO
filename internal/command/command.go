package command

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

func Execute(datastore *Datastore, input []string) error {
	switch input[0] {
	case Write:
		WriteData(GetLastTransaction(), datastore, input[1], input[2])
	case Read:
		ReadAndPrintData(GetLastTransaction(), datastore, input[1])
	}

	return nil
}

package command

import (
	"errors"
	"fmt"
	"strings"
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

var CmdArgLength = map[string]int{
	Write:  2,
	Read:   1,
	Delete: 1,
	Start:  0,
	Commit: 0,
	Abort:  0,
	Quit:   0,
}

// ValidateInput validates the command and the argument length
func ValidateInput(input []string) error {
	if len(input) == 0 {
		return errors.New("empty input")
	}

	cmd := strings.ToLower(input[0])
	inputArgLength := len(input) - 1

	argLength, ok := CmdArgLength[cmd]
	if !ok {
		return fmt.Errorf("invalid command: %s", cmd)
	}

	if argLength != inputArgLength {
		return fmt.Errorf("invalid number of arguments: %d", inputArgLength)
	}

	return nil
}

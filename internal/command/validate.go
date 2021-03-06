package command

import (
	"errors"
	"fmt"
)

// CmdArgLength represents the command and its respective argument length mapping
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
	if len(input) == 0 || input[0] == "" {
		return errors.New("empty input")
	}

	cmd := input[0]
	inputArgLength := len(input) - 1

	argLength, ok := CmdArgLength[cmd]
	if !ok {
		return fmt.Errorf("invalid command: %s", cmd)
	}

	if argLength != inputArgLength {
		return fmt.Errorf("invalid number of arguments: %d, needed: %d", inputArgLength, argLength)
	}

	return nil
}

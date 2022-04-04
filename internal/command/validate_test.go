package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateInputNoError tests validate input when there is no error
func TestValidateInputNoError(t *testing.T) {
	check := assert.New(t)
	input := []string{"write", "a", "hello"}

	err := ValidateInput(input)
	check.Nil(err)
}

// TestValidateInputEmptyInput tests validate input when input is empty
func TestValidateInputEmptyInput(t *testing.T) {
	check := assert.New(t)
	input := []string{""}

	err := ValidateInput(input)
	check.Contains(err.Error(), "empty input")
}

// TestValidateInputInvalidCommand tests validate input when invalid command is passed
func TestValidateInputInvalidCommand(t *testing.T) {
	check := assert.New(t)
	input := []string{"exit"}

	err := ValidateInput(input)
	check.Contains(err.Error(), "invalid command")
}

// TestValidateInputInvalidArgs tests validate input when invalid args are passed
func TestValidateInputInvalidArgs(t *testing.T) {
	check := assert.New(t)
	input := []string{"read"}

	err := ValidateInput(input)
	check.Contains(err.Error(), "invalid number of arguments")
}

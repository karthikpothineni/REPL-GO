package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCopyMap tests if input and copied map is having same data
func TestCopyMap(t *testing.T) {
	check := assert.New(t)
	input := map[string]string{
		"a": "hello",
		"b": "hi",
	}

	copiedMap := copyMap(input)
	check.EqualValues(copiedMap, input)

	// updating copied map should not change input map
	copiedMap["a"] = "hi"
	check.NotEqualValues(copiedMap, input)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"REPL-GO/internal/command"
)

// printRepl prints Repl prefix
func printRepl() {
	fmt.Print("> ")
}

// readInput reads the input from stdin
func readInput(reader *bufio.Reader) ([]string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	inputList := strings.Split(strings.TrimSuffix(input, "\n"), " ")

	// convert command to lowercase
	if len(inputList) > 0 {
		inputList[0] = strings.ToLower(inputList[0])
	}

	return inputList, nil
}

func main() {
	datastore := command.NewDataStore()
	reader := bufio.NewReader(os.Stdin)

	for {
		printRepl()
		inputList, err := readInput(reader)
		if err != nil {
			fmt.Println("Error in reading input", err.Error())
			continue
		}

		err = command.ValidateInput(inputList)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		command.Execute(datastore, inputList)
	}
}

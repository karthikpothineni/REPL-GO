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

// canContinue checks if repl can be continued
func canContinue(cmd string) bool {
	if strings.EqualFold(command.Quit, cmd) {
		return false
	}
	return true
}

// readInput reads the input from stdin
func readInput(reader *bufio.Reader) ([]string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	inputList := strings.Split(strings.TrimSuffix(input, "\n"), " ")
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

		if !canContinue(inputList[0]) {
			fmt.Println("Exiting...")
			return
		}
	}
}

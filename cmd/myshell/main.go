package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	input string
	args  []string
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		clean()

		cmd := getCmd(reader)

		switch cmd.input {
		case "exit":
			if cmd.args[0] == "0" {
				os.Exit(0)
			}
		case "echo":
			output(strings.Join(cmd.args, " "))
		default:
			output(fmt.Sprintf("%s: command not found", cmd.input))
		}
	}
}

func output(message string) {
	fmt.Fprint(os.Stdout, message+"\n")
}

func clean() {
	fmt.Fprint(os.Stdout, "$ ")
}

func getCmd(reader *bufio.Reader) *Command {

	// Wait for user input
	input, _ := reader.ReadString('\n')

	// Rimuove lo spazio finale
	arguments := strings.TrimSuffix(input, "\n")
	splits := strings.Split(arguments, " ")

	cmd := splits[0]

	args := splits[1:]

	return &Command{
		cmd,
		args,
	}
}

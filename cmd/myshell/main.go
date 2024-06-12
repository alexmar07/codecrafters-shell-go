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

		cmd.exec()
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

func (cmd *Command) exec() {

	switch cmd.input {
	case "exit":
		if cmd.args[0] == "0" {
			os.Exit(0)
		}
	case "echo":
		output(strings.Join(cmd.args, " "))
	case "type":
		output(typedCmd(cmd.args[0]))
	default:
		output(fmt.Sprintf("%s: command not found", cmd.input))
	}
}

func typedCmd(cmd string) string {

	switch cmd {
	case "echo":
		return "echo is a shell builtin"
	case "exit":
		return "exit is a shell builtin"
	case "type":
		return "type is a shell builtin"
	}

	return fmt.Sprintf("%s: not found", cmd)
}

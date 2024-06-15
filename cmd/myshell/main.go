package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/internal/commands"
)

type Command struct {
	input string
	args  []string
}

func main() {

	k := commands.NewKernel()

	reader := bufio.NewReader(os.Stdin)

	for {

		clean()

		cmd := getCmd(reader)

		fn, err := k.GetFn(cmd.input)

		if err != nil {

			if _, ok := err.(*commands.NotFoundCmdError); ok {

				if k.IsExternalCmd(cmd.input) {
					k.Exec(cmd.input, cmd.args)
					continue
				}

				output(err.Error())
			}

			continue
		}

		fn(cmd.args)
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

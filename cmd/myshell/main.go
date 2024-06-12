package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		clean()

		cmd := getCmd(reader)

		switch cmd {
		default:
			output(fmt.Sprintf("%s: command not found\n", cmd))
		}
	}
}

func output(message string) {
	fmt.Fprint(os.Stdout, message)
}

func clean() {
	fmt.Fprint(os.Stdout, "$ ")
}

func getCmd(reader *bufio.Reader) string {

	// Wait for user input
	input, _ := reader.ReadString('\n')

	// Rimuove lo spazio finale
	return strings.TrimSuffix(input, "\n")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// Rimuove lo spazio finale
	command := strings.TrimSuffix(input, "\n")

	switch command {
	default:
		output(fmt.Sprintf("%s: command not found\n", command))
	}
}

func output(message string) {
	fmt.Fprint(os.Stdout, message)
}

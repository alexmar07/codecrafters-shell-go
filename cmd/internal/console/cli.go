package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	cmd  string
	args []string
}

func NewCLI() *CLI {
	return &CLI{}
}

func (c CLI) GetCmd() string {
	return c.cmd
}

func (c CLI) GetArgs() []string {
	return c.args
}

func (c *CLI) GetInput() {

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSuffix(input, "\n")

	splits := strings.Split(input, " ")

	c.cmd = splits[0]
	c.args = splits[1:]

}

func (c *CLI) Clean() {
	fmt.Fprint(os.Stdout, "$ ")
}

func (c *CLI) Output(message string) {
	fmt.Fprint(os.Stdout, message+"\n")
}

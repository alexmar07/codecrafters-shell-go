package commands

import "fmt"

type cmdFnc func([]string)

type NotFoundCmdError struct {
	cmd string
}

func (e *NotFoundCmdError) Error() string {
	return fmt.Sprintf("%s: command not found", e.cmd)
}

type Kernel struct {
	commands map[string]cmdFnc
}

func NewKernel() *Kernel {

	k := &Kernel{make(map[string]cmdFnc)}

	k.Boot()

	return k
}

func (k *Kernel) registerCommand(cmd string, fn cmdFnc) {
	k.commands[cmd] = fn
}

func (k *Kernel) existsCommand(cmd string) bool {

	_, ok := k.commands[cmd]

	return ok
}

func (k *Kernel) Exec(cmd string) (cmdFnc, error) {

	if !k.existsCommand(cmd) {
		return nil, &NotFoundCmdError{cmd}
	}

	return k.commands[cmd], nil
}

func (k *Kernel) Boot() {
	k.registerCommand("exit", exit)
	k.registerCommand("echo", echo)
	k.registerCommand("type", k.typer)
}

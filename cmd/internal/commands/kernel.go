package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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

func (k *Kernel) isBuiltinCmd(cmd string) bool {

	_, ok := k.commands[cmd]

	return ok
}

func (k *Kernel) Exec(cmd string, args []string) {
	extCmd := exec.Command(cmd, args...)

	res, _ := extCmd.Output()
	fmt.Print(string(res))
}

func (k *Kernel) IsExternalCmd(cmd string) bool {

	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {

		// Create paths with command
		fp := filepath.Join(path, cmd)

		// Search command
		if _, err := os.Stat(fp); err == nil {
			return true
		}
	}

	return false
}

func (k *Kernel) GetFn(cmd string) (cmdFnc, error) {

	if !k.isBuiltinCmd(cmd) {
		return nil, &NotFoundCmdError{cmd}
	}

	return k.commands[cmd], nil
}

func (k *Kernel) Boot() {
	k.registerCommand("exit", exit)
	k.registerCommand("echo", echo)
	k.registerCommand("type", k.typer)
}

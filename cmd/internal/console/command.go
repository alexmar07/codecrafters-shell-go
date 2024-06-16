package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func exit(args []string) {

	if len(args) == 0 {
		os.Exit(1)
	}

	if code, err := strconv.Atoi(args[0]); err == nil {
		os.Exit(code)
	}
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func (k *Kernel) typer(args []string) {

	if len(args) == 0 {
		os.Exit(1)
	}

	if k.isBuiltinCmd(args[0]) {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return
	}

	// Retrieve all paths
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {

		// Create paths with command
		fp := filepath.Join(path, args[0])

		// Search command
		if _, err := os.Stat(fp); err == nil {
			fmt.Printf("%s is %s\n", args[0], fp)
			return
		}
	}

	fmt.Printf("%s: not found\n", args[0])
}

func cd(args []string) {

	homeDir, _ := os.UserHomeDir()

	if len(args) == 0 {
		os.Chdir(homeDir)
		return
	}

	if args[0] == "~" {
		os.Chdir(homeDir)
		return
	}

	var path string

	if strings.HasPrefix(args[0], "~") {
		path = homeDir + strings.TrimPrefix(args[0], "~")
	} else {
		path = args[0]
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("cd: %s: No such file or directory\n", args[0])
		return
	}

	os.Chdir(path)
}

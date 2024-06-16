package main

import (
	"github.com/codecrafters-io/shell-starter-go/cmd/internal/console"
)

func main() {

	cli := console.NewCLI()
	k := console.NewKernel()

	for {

		cli.Clean()
		cli.GetInput()

		fn, err := k.GetFn(cli.GetCmd())

		if err == nil {
			fn(cli.GetArgs())
			continue
		}

		if _, ok := err.(*console.NotFoundCmdError); ok {

			if !k.IsExternalCmd(cli.GetCmd()) {
				cli.Output(err.Error())
				continue
			}

			k.Exec(cli.GetCmd(), cli.GetArgs())
		}

	}

}

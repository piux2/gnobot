package main

import (
	"os"

	"github.com/gnolang/gno/pkgs/command"
	"github.com/gnolang/gno/pkgs/errors"
)

type (
	// AppItem is an app mapped to a subcommand
	AppItem = command.AppItem

	// AppList holds a list of subcommands apps
	AppList = command.AppList
)

var mainApps AppList = []AppItem{
	{App: faucetApp, Name: "faucet", Desc: "discord faucet", Defaults: DefaultFaucetOptions},
}

func main() {
	cmd := command.NewStdCommand()
	exec := os.Args[0]
	args := os.Args[1:]
	err := runMain(cmd, exec, args)
	if err != nil {
		cmd.ErrPrintfln("%s", err.Error())
		cmd.ErrPrintfln("%#v", err)
		return // exit
	}
}

func runMain(cmd *command.Command, exec string, args []string) error {
	// show help message.
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" {
		cmd.Println("available subcommands:")
		for _, appItem := range mainApps {
			cmd.Printf("  %s - %s\n", appItem.Name, appItem.Desc)
		}
		return nil
	}

	// switch on first argument.
	for _, appItem := range mainApps {
		if appItem.Name == args[0] {
			err := cmd.Run(appItem.App, args[1:], appItem.Defaults)
			return err // done
		}
	}

	// unknown app command!
	return errors.New("unknown command " + args[0])
}

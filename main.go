package main

import (
	"context"
	_ "embed"
	"os"

	"github.com/fatih/color"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main () {
	subcommands := []*ffcli.Command{
		stringCommand(),
		bytesCommand(),
		fsCommand(),
	}

	root := &ffcli.Command{
		Exec: func(ctx context.Context, args []string) error {
			printSubcommands(subcommands)
			return nil
		},
		Subcommands: subcommands,
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		color.Red(err.Error())
	}
}
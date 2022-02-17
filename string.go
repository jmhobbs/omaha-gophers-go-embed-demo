package main

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
)

//go:embed notes/hello.txt
var helloString string

func stringCommand() *ffcli.Command {
	return &ffcli.Command{
		Name:       "string",
		ShortUsage: "demo string",
		ShortHelp:  "Use go:embed into a string type",
		Exec:       func(_ context.Context, args []string) error {
			section("A single file can be embedded into a string")
			printCode(11, 12)
			fmt.Printf(" Type: %T\n", helloString)
			fmt.Printf("Value: %v\n", helloString)
			return nil
		},
	}
}
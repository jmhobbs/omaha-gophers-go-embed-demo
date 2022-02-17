package main

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
)

//go:embed notes/hello.txt
var helloBytes []byte

func bytesCommand() *ffcli.Command {
	return &ffcli.Command{
		Name:       "bytes",
		ShortUsage: "demo bytes",
		ShortHelp:  "Use go:embed into a byte slice type",
		Exec:       func(_ context.Context, args []string) error {
			section("A single file can be embedded into a byte slice")
			printCode(11, 12)
			fmt.Printf("  Type: %T\n", helloBytes)
			fmt.Printf(" Value: %v\n", helloBytes)
			fmt.Printf("String: %s\n", helloBytes)
			return nil
		},
	}
}
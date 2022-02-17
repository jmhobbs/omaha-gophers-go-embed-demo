package main

import (
	"bufio"
	"embed"
	"fmt"
	"path"
	"runtime"

	"github.com/fatih/color"
	"github.com/peterbourgon/ff/v3/ffcli"
)

//go:embed *.go
var code embed.FS

func printCode(startLine int, stopLine int) error {
	_, file, _, _ := runtime.Caller(1)

	f, err := code.Open(path.Base(file))
	if err != nil {
		return err
	}
	defer f.Close()

	color.Blue("------------------------------------------------------------")
	scanner := bufio.NewScanner(f)
	currentLine := 0
	for scanner.Scan() {
		currentLine = currentLine + 1
		if currentLine > stopLine {
			break
		}
		if currentLine >= startLine {
			color.Magenta(scanner.Text())
		}
	}
	color.Blue("------------------------------------------------------------")
	return scanner.Err()
}

func section(msg string) {
	color.Red("============================================================")
	color.Yellow(msg)
}

func printSubcommands(subcommands []*ffcli.Command) {
	color.Yellow("Subcommands")
	color.Red("===========")
	for _, cmd := range subcommands {
		fmt.Printf("%s : %s\n", color.CyanString("% 10s", cmd.Name), cmd.ShortHelp)
	}
}
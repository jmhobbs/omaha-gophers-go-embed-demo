package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io/fs"

	"github.com/peterbourgon/ff/v3/ffcli"
)

//go:embed images/gopher.png
var gopher embed.FS

//go:embed meetups
var meetups embed.FS

//go:embed images/*.gif
var gifs embed.FS

//go:embed images/*.png
//go:embed notes
var stacked embed.FS

//go:embed images/*.png notes
var oneline embed.FS

//go:embed notes
var notes embed.FS

//go:embed notes/*
var secretsExposed embed.FS

//go:embed templates
var templates embed.FS

func fsCommand() *ffcli.Command {
	single := &ffcli.Command{
		Name:       "single",
		ShortUsage: "demo fs single",
		ShortHelp:  "Embed a single file into an embed.FS",
		Exec:       func(_ context.Context, args []string) error {
			section("A single file in embed.FS preserves it's path")
			printCode(13, 14)
			return fs.WalkDir(gopher, ".", printTree)
		},
	}

	directory := &ffcli.Command{
		Name:       "directory",
		ShortUsage: "demo fs directory",
		ShortHelp:  "Embed an entire directory",
		Exec:       func(_ context.Context, args []string) error {
			section("A directory in embed.FS preserves the nested content")
			printCode(16, 17)
			return fs.WalkDir(meetups, ".", printTree)
		},
	}

	wildcards := &ffcli.Command{
		Name:       "wildcards",
		ShortUsage: "demo fs wildcards",
		ShortHelp:  "Embed files with wildcard patterns",
		Exec:       func(_ context.Context, args []string) error {
			section("A wildcard pattern in embed.FS can be used to select specific files")
			printCode(19, 20)
			return fs.WalkDir(gifs, ".", printTree)
		},
	}

	stacking := &ffcli.Command{
		Name:       "stacking",
		ShortUsage: "demo fs stacking",
		ShortHelp:  "Use multiple directives on one embed.FS",
		Exec:       func(_ context.Context, args []string) error {
			section("A patterns can be stacked into a single embed.FS")
			printCode(22, 24)
			err := fs.WalkDir(stacked, ".", printTree)
			if err != nil {
				return err
			}
			section("These can also be combined into one line")
			printCode(26, 27)
			return fs.WalkDir(oneline, ".", printTree)
		},
	}

	exclusions := &ffcli.Command{
		Name:       "exclusions",
		ShortUsage: "demo fs exclusions",
		ShortHelp:  "Automatically exclude hidden files in directories",
		Exec:       func(_ context.Context, args []string) error {
			section("Directories ignore files starting in . or _")
			printCode(29, 30)
			err := fs.WalkDir(notes, ".", printTree)
			if err != nil {
				return err
			}
			section("To include these files, use a wildcard")
			printCode(33,34)
			return fs.WalkDir(secretsExposed, ".", printTree)
		},
	}

	templatesCmd := &ffcli.Command{
		Name:       "templates",
		ShortUsage: "demo fs templates",
		ShortHelp:  "Use go:embed into a embed.FS type",
		Exec:       func(_ context.Context, args []string) error {
			section("embed.FS can be used with anything that accepts fs.FS")
			printCode(35, 36)
			tmpls, err := template.ParseFS(templates, "**/*.html")
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Printf("%T\n", tmpls)
			for _, tmpl := range tmpls.Templates() {
				fmt.Println("-", tmpl.Name())
			}
			return nil
		},
	}

	subcommands := []*ffcli.Command{
		single,
		directory,
		wildcards,
		stacking,
		exclusions,
		templatesCmd,
	}

	return &ffcli.Command{
		Name:       "fs",
		ShortUsage: "demo fs",
		ShortHelp:  "Use go:embed into a embed.FS type",
		Exec:       func(_ context.Context, args []string) error {
			printSubcommands(subcommands)
			return nil
		},
		Subcommands: subcommands,
	}
}

func printTree(path string, d fs.DirEntry, err error) error {
	if path != "." {
		fmt.Printf("./%s\n", path)
	}
	return err
}
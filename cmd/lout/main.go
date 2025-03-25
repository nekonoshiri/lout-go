package main

import (
	"flag"
	"fmt"
	"os"
)

var file string

func init() {
	flag.StringVar(&file, "f", "", "")
	flag.StringVar(&file, "file", "", "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: lout -f <state file> command

A low-level runtime for "Lights Out"-like games.

Options:
  -h, --help     show help
  -f, --file     specify state file (required)

Commands:
  init <w> <h>   initialize state file with width w and height h
  press <x> <y>  press the light at (x, y)
`)
	}
}

func main() {
	flag.Parse()

	if file == "" {
		fmt.Fprintln(os.Stderr, "state file is required")
		flag.Usage()
		os.Exit(2)
	}

	switch cmd := flag.Arg(0); cmd {
	case "init":
		// TODO
	case "press":
		// TODO
	case "":
		fmt.Fprintln(os.Stderr, "command is required")
		flag.Usage()
		os.Exit(2)
	default:
		fmt.Fprintln(os.Stderr, "invalid command:", cmd)
		flag.Usage()
		os.Exit(2)
	}
}

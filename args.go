package main

import (
	"flag"
	"os"
	"strconv"
)

type Args struct {
	Title  string
	ID     int
	To     int
	Create bool
	List   bool
	Delete bool
	Move   bool
}

var a Args = Args{}

func ParseArgs() *Args {
	flag.Parse()

	if a.Create {
		parseCreateArgs()
	}

	if a.Delete {
		parseDeleteArgs()
	}

	if a.Move {
		parseMoveArgs()
	}

	return &a
}

func parseCreateArgs() {
	if flag.NArg() < 1 {
		usageAndExit()
	}

	a.Title = flag.Arg(0)
}

func parseDeleteArgs() {
	a.ID = parseOptionalInt(0, 0)
}

func parseMoveArgs() {
	a.ID = parseInt(0)
	a.To = parseOptionalInt(1, 0)
}

func parseInt(arg int) int {
	if flag.NArg() < arg {
		usageAndExit()
	}

	a := flag.Arg(arg)
	i, err := strconv.Atoi(a)
	if err != nil {
		usageAndExit()
	}

	return i
}

func parseOptionalInt(arg int, fallback int) int {
	if flag.NArg() < arg {
		return fallback
	}

	a := flag.Arg(arg)
	if a == "" {
		return fallback
	}

	i, err := strconv.Atoi(a)
	if err != nil {
		usageAndExit()
	}

	return i
}

func usageAndExit() {
	flag.Usage()
	os.Exit(1)
}

func init() {
	flag.BoolVar(&a.Create, "create", false, "Create a new todo")
	flag.BoolVar(&a.List, "list", false, "List all todos")
	flag.BoolVar(&a.Delete, "delete", false, "Delete a todo")
	flag.BoolVar(&a.Move, "move", false, "Move a todo")
}

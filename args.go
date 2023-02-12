package main

import (
	"flag"
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
		flag.Usage()
	}

	a.Title = flag.Arg(0)
}

func parseDeleteArgs() {
	if flag.NArg() < 1 {
		flag.Usage()
	}

	i := flag.Arg(0)
	id, err := strconv.Atoi(i)
	if err != nil {
		flag.Usage()
	}

	a.ID = id
}

func parseMoveArgs() {
	if flag.NArg() < 2 {
		flag.Usage()
	}

	f := flag.Arg(0)
	from, err := strconv.Atoi(f)
	if err != nil {
		flag.Usage()
	}

	t := flag.Arg(1)
	to, err := strconv.Atoi(t)
	if err != nil {
		flag.Usage()
	}

	a.ID = from
	a.To = to
}

func init() {
	flag.BoolVar(&a.Create, "create", false, "Create a new todo")
	flag.BoolVar(&a.List, "list", false, "List all todos")
	flag.BoolVar(&a.Delete, "delete", false, "Delete a todo")
	flag.BoolVar(&a.Move, "move", false, "Move a todo")
}

package main

import "fmt"

var (
	// Version is the version of the application
	Version = "0.0.0"
)

func main() {
	c := NewContext()

	if c.Args.Create {
		createTodo(c)
	}

	if c.Args.List {
		listTodos(c)
	}

	if c.Args.Delete {
		deleteTodo(c)
	}

	if c.Args.Move {
		moveTodo(c)
	}
}

func createTodo(c *Context) {
	c.Store.Add(Todo{Title: c.Args.Title})
	err := c.Store.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func listTodos(c *Context) {
	for i, t := range c.Store.data {
		fmt.Printf("%d: %s\n", i, t.Title)
	}
}

func deleteTodo(c *Context) {
	c.Store.Delete(c.Args.ID)
	err := c.Store.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func moveTodo(c *Context) {
	c.Store.Move(c.Args.ID, c.Args.To)
	err := c.Store.Save()
	if err != nil {
		fmt.Println(err)
	}
}

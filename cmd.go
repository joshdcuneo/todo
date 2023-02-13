package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "todo",
		Short: "A simple todo app",
		Long:  `A simple todo app`,
	}
}

func createCmd(c *Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new todo",
		Long:  `Create a new todo`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c.Store.Add(Todo{Title: strings.TrimSpace(args[0])})
			err := c.Store.Save()
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	return cmd
}

func listCmd(c *Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all todos",
		Long:  `List all todos`,
		Run: func(cmd *cobra.Command, args []string) {
			for i, t := range c.Store.data {
				fmt.Printf("%d: %s\n", i, t.Title)
			}
		},
	}
}

func deleteCmd(c *Context) *cobra.Command {
	var todo int

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a todo",
		Long:  "Delete a todo",
		Run: func(cmd *cobra.Command, args []string) {
			c.Store.Delete(todo)
			err := c.Store.Save()
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	cmd.Flags().IntVarP(&todo, "todo", "t", 0, "ID of the todo")

	return cmd
}

func moveCmd(c *Context) *cobra.Command {
	var todo, move int

	cmd := &cobra.Command{
		Use:   "move",
		Short: "Move a todo",
		Long:  "Move a todo",
		Run: func(cmd *cobra.Command, args []string) {
			c.Store.Move(todo, move)
			err := c.Store.Save()
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	cmd.Flags().IntVarP(&todo, "todo", "t", 0, "ID of the todo")
	cmd.MarkFlagRequired("todo")
	cmd.Flags().IntVarP(&move, "move", "m", 0, "ID to move the todo to")

	return cmd
}

func Execute(c *Context) {
	r := rootCmd()
	r.AddCommand(createCmd(c))
	r.AddCommand(listCmd(c))
	r.AddCommand(deleteCmd(c))
	r.AddCommand(moveCmd(c))

	if err := r.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

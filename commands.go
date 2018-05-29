package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// GlobalFlags is cli flags
var GlobalFlags = []cli.Flag{}

// Action is main function
func Action(c *cli.Context) error {
	fmt.Println("hello")
	return nil
}

// CommandNotFound is exception function
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

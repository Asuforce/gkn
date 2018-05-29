package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// GlobalFlags is cli flags
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "file, f",
		Value: "",
		Usage: "perse from file",
	},
}

// Action is main function
func Action(c *cli.Context) error {
	if c.String("f") != "" {
		_, e := ioutil.ReadFile(c.String("f"))
		if e != nil {
			fmt.Println(e)
			return nil
		}
	}

	a := strings.Split(c.Args()[0], " ")
	s := []string{}

	s = append(s, a[0])
	for i := range a {
		c := 0
		for j := range s {
			if a[i] == s[j] {
				break
			}
			c++

			if c == len(s) {
				s = append(s, a[i])
			}
		}
	}

	for i := range s {
		fmt.Println(s[i])
	}

	fmt.Println(len(s))

	return nil
}

// CommandNotFound is exception function
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

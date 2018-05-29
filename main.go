package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "asuforce"
	app.Usage = "count paragraph from your text"
	app.Flags = GlobalFlags
	app.Action = Action
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}

package main

import (
	"os"

	"github.com/midN/awtf/commands"
	"github.com/midN/awtf/flags"
	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Name = "AWTF"
	app.Usage = "AWS CLI Wrapper for complicated aws cli commands"

	app.Flags = flags.Flags()
	app.Commands = commands.Commands()

	app.Run(os.Args)
}

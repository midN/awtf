package main

import (
	"os"

	cmds "github.com/midN/awtf/commands"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "AWTF"
	app.Usage = "AWS CLI Wrapper for complicated aws cli commands"

	app.Commands = []cli.Command{
		cmds.CfCommands(),
	}

	app.Run(os.Args)
}

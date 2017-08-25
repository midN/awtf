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

	app.Flags = globalFlags()
	app.Commands = cmds.Commands()

	app.Run(os.Args)
}

func globalFlags() []cli.Flag {
	var profile string
	var region string

	profileFlag := cli.StringFlag{
		Name:        "profile, p",
		Usage:       "Optional AWS profile",
		Value:       "default",
		Destination: &profile,
	}

	regionFlag := cli.StringFlag{
		Name:        "region, r",
		Usage:       "Optional AWS region",
		Value:       "eu-west-1",
		Destination: &region,
	}

	return []cli.Flag{
		profileFlag,
		regionFlag,
	}
}

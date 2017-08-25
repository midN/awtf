package commands

import (
	"github.com/fatih/color"
	cf "github.com/midN/awtf/cloudformation"
	"gopkg.in/urfave/cli.v1"
)

var (
	// Vars for flags
	stackName string

	// SubCommands
	lpCommand = cli.Command{
		Name:    "list-params",
		Aliases: []string{"lp"},
		Usage:   "List existing cf stack params",
		Action: func(c *cli.Context) error {
			err := cf.ListParams(stackName)

			if err != nil {
				redError := color.RedString(err.Error())
				return cli.NewExitError(redError, 1)
			} else {
				return nil
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "stack-name, sn",
				Usage:       "existing stacks name",
				Destination: &stackName,
			},
		},
	}

	// Commands
	cfCommand = cli.Command{
		Name:    "cloudformation",
		Aliases: []string{"cf"},
		Usage:   "cloudformation related commands",
		Subcommands: []cli.Command{
			lpCommand,
		},
	}
)

func CfCommands() cli.Command {
	return cfCommand
}

package commands

import (
	"github.com/fatih/color"
	cf "github.com/midN/awtf/cloudformation"
	"github.com/midN/awtf/common"
	"gopkg.in/urfave/cli.v1"
)

var (
	// Vars for flags
	stackName string

	// Flags
	snFlag = cli.StringFlag{
		Name:        "stack-name, sn",
		Usage:       "existing stacks name",
		Destination: &stackName,
	}

	// Actions
	lpAction = func(c *cli.Context) error {
		lpParams := cf.LpParams{
			StackName: stackName,
			Sess:      common.InitSession(c),
		}

		err := cf.ListParams(lpParams)

		if err != nil {
			redError := color.RedString(err.Error())
			return cli.NewExitError(redError, 1)
		} else {
			return nil
		}
	}

	// SubCommands
	lpCommand = cli.Command{
		Name:    "list-params",
		Aliases: []string{"lp"},
		Usage:   "List existing cf stack params",
		Action:  lpAction,
		Flags: []cli.Flag{
			snFlag,
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

func cfCommands() cli.Command {
	return cfCommand
}

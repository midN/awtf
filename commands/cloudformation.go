package commands

import (
	"github.com/midN/awtf/actions"
	"github.com/midN/awtf/flags"
	"gopkg.in/urfave/cli.v1"
)

var (
	// SubCommands
	lpCommand = cli.Command{
		Name:    "list-params",
		Aliases: []string{"lp"},
		Usage:   "List existing cf stack params",
		Action:  actions.CfLpAction(),
		Flags:   flags.CfLpFlags(),
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

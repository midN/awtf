package commands

import (
	"gopkg.in/urfave/cli.v1"
)

func Commands() []cli.Command {
	return []cli.Command{
		cfCommands(),
	}
}

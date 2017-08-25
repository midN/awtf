package flags

import "gopkg.in/urfave/cli.v1"

func CfLpFlags() []cli.Flag {
	// Flags
	snFlag := cli.StringFlag{
		Name:  "stack-name, sn",
		Usage: "existing stacks name",
	}

	return []cli.Flag{
		snFlag,
	}
}

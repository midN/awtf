package flags

import "gopkg.in/urfave/cli.v1"

func CfLpFlags() []cli.Flag {
	// Flags
	snFlag := cli.StringFlag{
		Name:  "stack-name, sn",
		Usage: "existing stacks name",
	}

	ttFlag := cli.StringFlag{
		Name:  "table-type, tt",
		Usage: "table type, Options available: ( markdown, default )",
		Value: "default",
	}

	return []cli.Flag{
		snFlag,
		ttFlag,
	}
}

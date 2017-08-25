package flags

import "gopkg.in/urfave/cli.v1"

func GlobalFlags() []cli.Flag {
	profileFlag := cli.StringFlag{
		Name:  "profile, p",
		Usage: "Optional AWS profile",
		Value: "default",
	}

	regionFlag := cli.StringFlag{
		Name:  "region, r",
		Usage: "Optional AWS region",
		Value: "eu-west-1",
	}

	return []cli.Flag{
		profileFlag,
		regionFlag,
	}
}

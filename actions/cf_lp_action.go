package actions

import (
	"github.com/fatih/color"
	"github.com/midN/awtf/cloudformation"
	"github.com/midN/awtf/common"
	"gopkg.in/urfave/cli.v1"
)

func CfLpAction() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		lpParams := cloudformation.LpParams{
			StackName: c.String("stack-name"),
			Sess:      common.InitSession(c),
			TableType: c.String("table-type"),
		}

		err := cloudformation.ListParams(lpParams, c.App.Writer)

		if err != nil {
			redError := color.RedString(err.Error())
			return cli.NewExitError(redError, 1)
		} else {
			return nil
		}
	}
}

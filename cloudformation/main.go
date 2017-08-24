package cloudformation

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func Welcome() {
	fmt.Print("You are in CF package")
}

func PrintParams(stackName string) error {
	stackNameMissing := color.RedString("Stack name is missing!")

	if stackName != "" {
		fmt.Println("Your stack:", stackName)
		return nil
	} else {
		return cli.NewExitError(stackNameMissing, 1)
	}
}

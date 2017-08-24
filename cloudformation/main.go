package cloudformation

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func ListParams(stackName string) error {
	// TODO: Add wrapper function for Colorizing error-output, or maybe some other way?:)
	stackNameMissing := color.RedString("Stack name is missing!")

	if stackName == "" {
		return cli.NewExitError(stackNameMissing, 1)
	}

	// TODO: Move session to main or separate package?
	// TODO: Describe where is region taken from in AWS.
	// TODO: Add options to pass profile/region, example below
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	Profile: "X",
	//	Config:  aws.Config{Region: aws.String("eu-west-1")},
	//}))
	sess := session.Must(session.NewSession())

	cfSession := cloudformation.New(sess)
	params := cloudformation.DescribeStacksInput{StackName: aws.String(stackName)}

	stackInfo, err := cfSession.DescribeStacks(&params)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	// TODO: Print params with some nice table
	fmt.Println(stackInfo.Stacks[0].Parameters)

	return nil
}

package cloudformation

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func ListParams(stackName string) error {
	if stackName == "" {
		return errors.New("Stack name is missing!")
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
		return err
	}

	// TODO: Print params with some nice table
	fmt.Println(stackInfo.Stacks[0].Parameters)

	return nil
}

package cloudformation

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type LpParams struct {
	StackName string
	Sess      *session.Session
}

func ListParams(lpParams LpParams) error {
	if lpParams.StackName == "" {
		return errors.New("Stack name is missing!")
	}

	cfSession := cloudformation.New(lpParams.Sess)
	params := cloudformation.DescribeStacksInput{StackName: aws.String(lpParams.StackName)}
	stackInfo, err := cfSession.DescribeStacks(&params)

	if err != nil {
		return err
	}

	// TODO: Print params with some nice table
	fmt.Println(stackInfo.Stacks[0].Parameters)

	return nil
}

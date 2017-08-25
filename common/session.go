package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"gopkg.in/urfave/cli.v1"
)

func InitSession(c *cli.Context) *session.Session {
	profile := c.GlobalString("profile")
	region := c.GlobalString("region")

	return session.Must(
		session.NewSessionWithOptions(
			session.Options{
				Profile: profile,
				Config:  aws.Config{Region: aws.String(region)},
			}))
}

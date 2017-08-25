# AWTF Description
Wrapper for AWS CLI.

Have you ever tried to update existing CloudFormation stack with couple of params?

What if your stack params didn't change and you want to update with existing values ( Nested stacks )?

The result will be an extremely long aws cli command with every Parameter listed out.

That's why i'm writing this right-now-simple Go wrapper for AWS CLI starting with everyday CloudFormation commands i personally use.

## Authorization
Authorization is currently done based on default AWS authentication.
Please take a look at AWS documentation on setting up credentials, there are multiple different ways to do so.

**[AWS Credentials docs](http://docs.aws.amazon.com/cli/latest/topic/config-vars.html)**

**NB!** Please note that `region` is not taken from Credentials file for some reason.

Therefore for region to correctly work with aws-sdk-go, for now please use `AWS_REGION` env variable.

## Backlog / Ideas
Work in Progress and future ideas can be found at [GitHub Project](https://github.com/midN/awtf/projects)

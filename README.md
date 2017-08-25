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

**NB!** Default region is `eu-west-1`

`awtf` now supports additional global flags for region and profile.

## Commands
### List Params
List params of existing CloudFormation Stack.

**Example:** `awtf cf lp --stack-name test`

**Output:**

|   Parameter Name   |     Parameter Value      |
|--------------------|--------------------------|
| Hello              | X                        |
| Hello              | X                        |

## Backlog / Ideas
Work in Progress and future ideas can be found at [GitHub Project](https://github.com/midN/awtf/projects)

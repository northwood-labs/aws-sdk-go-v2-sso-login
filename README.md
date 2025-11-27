# AWS Identity Center (née SSO) login functionality for AWS SDK for Go v2

> [!WARNING]
> THIS IS NOT AN OFFICIAL PART OF `aws-sdk-go-v2`. This has not been created, endorsed, or validated by AWS.

## Why This Package Exists

The official [github.com/aws/aws-sdk-go-v2/service/sso](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sso) package does not contain a `login` function. AWS recommends using the AWS CLI `aws sso login` which requires that the user has the AWS CLI installed on their system.

## Goal of Package

This package’s goal is to remove the need for a user to have anything installed when using AWS SSO, and enabling developers to create tools that do not need to rely on the AWS CLI.

## Prerequisites

The user must have a `~/.aws/config` file with at least one `[profile <profileName>]` section in it. Read “[Configuring IAM Identity Center authentication with the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html)” for more information.

## References

This package was inspired by, and aims to solve this GH issue: https://github.com/aws/aws-sdk-go-v2/issues/1222

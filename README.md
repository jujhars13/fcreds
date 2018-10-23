# fcreds awssecrets

Fcreds awssecrets is a command line tool to interact with [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/).

awssecrets is written in go and was inspired by https://github.com/Versent/unicreds

## Usage

```bash
usage: awssecrets --region=REGION [<flags>] <command> [<args> ...]

A CLI tool to get secrets from AWS secrets manager.

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
  -r, --region=REGION    Configure the AWS region
  -p, --profile=PROFILE  Configure the AWS profile

Commands:
  help [<command>...]
    Show help.

  create <name> <value> [<description>]
    Create a secret

  exec --secret-name=SECRET-NAME <command>...
    Execute a command with all secrets loaded as environment variables.
```

awssecrets supports the AWS_* environment variables, and configuration in `~/.aws/`credentials` and `~/.aws/config`

## Example

1. Create a secret in AWS using the CLI

```bash
aws secretsmanager create-secret --name secret-friendly-name --secret-string secret-string-value
```

2. Execute `env` command, all secrets are loaded as environment variables.

```bash
awssecrets -r us-east-1 exec -n secret-name-1 -n secret-name-2 -- env
```

## Release

To release a new version you'll need Docker running on your machine and the environment variable GITHUB_TOKEN set locally. Then we can run

```bash
./release.sh v1.2
```

release.sh takes a version number as a parameter (or we'll try to release v1.0 by default)

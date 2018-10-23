# awssecrets

awssecrets is a command line tool to interact with AWS Secrets Manager.
awssecrets is written in go and was inspired by https://github.com/Versent/unicreds

# usage

```
usage: awssecrets --region=REGION [<flags>] <command> [<args> ...]

A command line tool to get AWS secrets.

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
  -r, --region=REGION    Configure the AWS region
  -p, --profile=PROFILE  Configure the AWS profile

Commands:
  help [<command>...]
    Show help.

  exec --secret-name=SECRET-NAME <command>...
    Execute a command with all secrets loaded as environment variables.
```

awssecrets supports the AWS_* environment variables, and configuration in ~/.aws/credentials and ~/.aws/config

# example

* Create a secret in AWS using the CLI
```
aws secretsmanager create-secret --name secret-friendly-name --secret-string secret-string-value
```
* Execute `env` command, all secrets are loaded as environment variables.
```
awssecrets -r us-east-1 exec -n secret-name-1 -n secret-name-2 -- env
```

## IAM policy

To be able to access the secret value, any role used to call fcreds will require an IAM policy that looks like this
```
- PolicyName: SecretsManagerAccess
  PolicyDocument:
    Statement:
    - Effect: Allow
      Action:
      - secretsmanager:GetSecretValue
      Resource:
      - arn:aws:secretsmanager:region:accountId:secret:secretsName-*
```

## Release

To release a new version you'll need Docker running on your machine and the environment variable GITHUB_TOKEN set locally. Then we can run
```
./release.sh v1.2
```
release.sh takes a version number as a parameter (or we'll try to release v1.0 by default)

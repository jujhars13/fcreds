package awssecrets

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type secretsManager interface {
	GetSecretValue(*secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
}

// Creates AWS session
func getSession(profile *string, region *string) *session.Session {
	config := aws.Config{Region: region}
	options := session.Options{
		Config:  config,
		Profile: aws.StringValue(profile),
	}

	sess, err := session.NewSessionWithOptions(options)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create AWS session: %v\n", err)
		os.Exit(1)
	}

	return sess
}

// GetSecretManager creates and return an aws secrets manager
func GetSecretManager(profile, region *string) *secretsmanager.SecretsManager {
	sess := getSession(profile, region)
	return secretsmanager.New(sess)
}

// Gets secret from AWS
func getSecret(svc secretsManager, name string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}

	result, err := svc.GetSecretValue(input)

	if err != nil {
		handleSecretError(err)

		return "", err
	}

	return *result.SecretString, nil
}

// Displays error returned when trying to interact with AWS secrets
func handleSecretError(err error) {
	// Print the error, cast err to awserr.Error to get the Code and Message from an error.
	if aerr, hasAwsError := err.(awserr.Error); hasAwsError {
		fmt.Println(aerr.Code(), aerr.Error())
	} else {
		fmt.Println(err.Error())
	}
}

// SetSecretEnv sets all requested secrets values as environment variables
func SetSecretEnv(svc secretsManager, secretList []string) {
	for _, secretName := range secretList {
		secret, err := getSecret(svc, secretName)

		if err == nil {
			err := os.Setenv(secretName, secret)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

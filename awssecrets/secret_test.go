package awssecrets

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/stretchr/testify/assert"
)

type secretsManagerStub struct{}

func (s secretsManagerStub) GetSecretValue(*secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	secretValue := "testSecret"
	return &secretsmanager.GetSecretValueOutput{SecretString: &secretValue}, nil
}

func TestGetSecret(t *testing.T) {
	secretManagerService := secretsManagerStub{}
	result, err := getSecret(secretManagerService, "test")

	assert.Nil(t, err)
	assert.Equal(t, "testSecret", result)
}

func TestSetEnv(t *testing.T) {
	secretManagerService := secretsManagerStub{}
	SetSecretEnv(secretManagerService, []string{"secret1", "secret2"})
	assert.Equal(t, "testSecret", os.Getenv("secret1"))
	assert.Equal(t, "testSecret", os.Getenv("secret2"))
}

func TestGetSession(t *testing.T) {
	profile := ""
	region := "testRegion"
	sess := getSession(&profile, &region)
	assert.Equal(t, "testRegion", *sess.Config.Region)
}

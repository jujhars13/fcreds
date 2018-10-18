package awssecrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetSecretList(t *testing.T) {
	array := new([]string)
	secretList := secretListValue(*array)
	result := secretList.Set("testString")

	assert.Nil(t, result)
	assert.Equal(t, 1, len(secretList))
	assert.Equal(t, "testString", secretList[0])
}

func TestStringSecretList(t *testing.T) {
	array := new([]string)
	secretList := secretListValue(*array)
	result := secretList.String()

	assert.Equal(t, "", result)
}

func TestIsCumulativeSecretList(t *testing.T) {
	array := new([]string)
	secretList := secretListValue(*array)
	result := secretList.IsCumulative()

	assert.Equal(t, true, result)
}

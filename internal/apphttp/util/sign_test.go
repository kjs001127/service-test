package util

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSigningKey(t *testing.T) {
	signingKey, signingKeyErr := CreateSigningKey()
	hexBytes, hexErr := hex.DecodeString(signingKey)

	assert.Nil(t, signingKeyErr)
	assert.Nil(t, hexErr)
	assert.Equal(t, 64, len(signingKey))
	assert.Equal(t, 32, len(hexBytes))
}

func TestSign(t *testing.T) {
	signingKey := "963b43f7bf0dd6ee4a111cbaefd425f217a293e8f9ac1c4ecd7eea0d9c420bc6"
	body := json.RawMessage(`{"method": "testMethod", "context": {"caller": {"type": "manager", "id": "123"}, "channel": {"id": "321"}}}`)
	sign, err := Sign(signingKey, body)

	assert.Nil(t, err)
	assert.Equal(t, "suJvM3E5YL4pIHTuMbMqKt7QsnCPLjXVSd7LrNlEA54=", sign)
}

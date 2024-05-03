package util

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func CreateSigningKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	signingKey := hex.EncodeToString(bytes)
	return signingKey, nil
}

func Sign(signingKey string, body []byte) (string, error) {
	hexBytes, hexErr := hex.DecodeString(signingKey)
	if hexErr != nil {
		return "", hexErr
	}
	hash := hmac.New(sha256.New, hexBytes)
	_, hmacErr := hash.Write(body)
	if hmacErr != nil {
		return "", hmacErr
	}
	s := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(s), nil
}

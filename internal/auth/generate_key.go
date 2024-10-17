package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSecureKey() string {
	secretBytes := make([]byte, 32)
	_, err := rand.Read(secretBytes)

	if err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(secretBytes)
}

package util

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
)

// GetAPIKeyHash - Get API Key hash
func GetAPIKeyHash(key string) string {
	return keyHash(key)
}

// NewAPIKey - Generates new API Key for client
func NewAPIKey() (string, string, error) {
	key, err := generateRandomString(64)

	if err != nil {
		return "", "", err
	}

	return key, keyHash(key), nil
}

func keyHash(key string) string {
	hasher := crypto.SHA512.New()
	hasher.Write([]byte(key))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

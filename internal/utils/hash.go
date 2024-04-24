package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateHash(s string) string {
	hash := sha256.Sum256([]byte(s))
	encodedHash := base64.URLEncoding.EncodeToString(hash[:])
	return encodedHash[:8]
}

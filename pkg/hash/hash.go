package hash

import (
	"crypto/sha1"
	"encoding/base64"
)

// Hash - function for hashing text by SHA-1 algorithm
func Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

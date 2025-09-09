package signutils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateHmacSHA256 Helper function to generate HMAC-SHA256 signature
func GenerateHmacSHA256(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

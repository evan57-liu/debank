package cacheutils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/coin50etf/coin-market/internal/pkg/json"
)

func generateCacheKey(any any) string {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		log.Println("failed to marshal cache key: ", err)
		return ""
	}

	hash := sha256.Sum256(jsonBytes)

	return hex.EncodeToString(hash[:])
}

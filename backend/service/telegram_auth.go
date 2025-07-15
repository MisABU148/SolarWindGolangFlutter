package service

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"sort"
	"strings"
)

func ValidateTelegramAuth(data map[string]string) bool {
	checkHash := data["hash"]
	delete(data, "hash")

	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var dataCheckString string
	for _, k := range keys {
		dataCheckString += k + "=" + data[k] + "\n"
	}
	dataCheckString = strings.TrimSuffix(dataCheckString, "\n")

	secretKey := sha256.Sum256([]byte(os.Getenv("TELEGRAM_BOT_TOKEN")))
	hash := sha256.Sum256(append(secretKey[:], []byte(dataCheckString)...))

	return hex.EncodeToString(hash[:]) == checkHash
}

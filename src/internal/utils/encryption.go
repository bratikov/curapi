package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"

	"github.com/zeebo/blake3"
)

func Encrypt(input string, key string) string {
	hash := blake3.Sum256([]byte(input + key))
	return hex.EncodeToString(hash[:])
}

func EncryptSha(input string, key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input + key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := len(charset)
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		randomBytes[i] = charset[int(randomBytes[i])%charsetLength]
	}

	return string(randomBytes)
}

func HashToInt(hash string) int {
	subset := (uint16(hash[0]) << 8) | uint16(hash[1])
	return int(subset)%15 + 1
}

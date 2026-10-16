package uuid

import (
	"fmt"
	"crypto/rand"
	"math/big"

	uuidpkg "github.com/google/uuid"
)

const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenUUID() (string, error) {
	id, err := uuidpkg.NewRandom()
	if err != nil {
		return "", fmt.Errorf("生成 uuid 失败: %w", err)
	}
	return id.String(), nil
}

func GenUUIDIfFailedGetEmpty() string {
	s, err := GenUUID()
	if err != nil {
		return ""
	}
	return s
}

// GenerateAlphaID generates a random alphabetic string of the given length.
func GenerateAlphaID(length int) (string, error) {
	b := make([]byte, length)
	alphaLen := big.NewInt(int64(len(alpha)))

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, alphaLen)
		if err != nil {
			return "", err
		}
		b[i] = alpha[num.Int64()]
	}

	return string(b), nil
}

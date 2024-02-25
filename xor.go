package cred

import (
	"strings"
)

func XorEncryptDecrypt(data, key string) string {
	var encryptedData []string
	keyLength := len(key)
	for i, char := range []rune(data) {
		keyChar := rune(key[i%keyLength])
		encryptedChar := string(char ^ keyChar)
		encryptedData = append(encryptedData, encryptedChar)
	}

	return strings.Join(encryptedData, "")
}
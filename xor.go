package cred

func XorEncryptDecrypt(data, key string) string {
	encryptedData := make([]byte, len(data))
	keyLength := len(key)

	for i, char := range data {
		keyChar := key[i%keyLength]
		encryptedChar := byte(char) ^ keyChar
		encryptedData[i] = encryptedChar
	}

	return string(encryptedData)
}
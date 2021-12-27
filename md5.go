package cred

import (
	"crypto/md5"
	"encoding/hex"
)

type Md5UserSaltCredManager struct{}

func getMd5(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}

func NewMd5() *Md5UserSaltCredManager {
	cm := &Md5UserSaltCredManager{}
	return cm
}

// GetMd5HexDigest Return a 32-bit md5 encrypted string
func (cm *Md5UserSaltCredManager) GetMd5HexDigest(plainText string) string {
	b := getMd5([]byte(plainText))
	return hex.EncodeToString(b)
}

// GetMd5HexDigest16 Return a 16-bit md5 encrypted string
func (cm *Md5UserSaltCredManager) GetMd5HexDigest16(plainText string) string {
	return cm.GetMd5HexDigest(plainText)[8:24]
}

func (cm *Md5UserSaltCredManager) IsMatch(plainText, hashedText string) bool {
	return hashedText == cm.GetMd5HexDigest(plainText)
}

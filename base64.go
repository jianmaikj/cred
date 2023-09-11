package cred

import "encoding/base64"

func Base64EncodeStr(data string) (res string) {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64EncodeByte(data []byte) (res string) {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) (res []byte, err error) {
	res, err = base64.StdEncoding.DecodeString(data)
	return
}
package cred

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"runtime"
)

type RSACredManager struct{}

func NewRSA() *RSACredManager {
	cm := &RSACredManager{}
	return cm
}

func (cm *RSACredManager)Encrypt(plainText ,key []byte)(cryptText []byte,err error){
	block, _:= pem.Decode(key)
	defer func(){
		if err:=recover();err!=nil{
			switch err.(type){
			case runtime.Error:
				log.Println("runtime err:",err,"Check that the key is correct")
			default:
				log.Println("error:",err)
			}
		}
	}()
	publicKeyInterface,err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		return nil,err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err!=nil{
		return nil,err
	}
	return cipherText,nil
}

func (cm *RSACredManager)Decrypt(cryptText ,key []byte)(plainText []byte,err error){
	block, _ := pem.Decode(key)

	defer func(){
		if err:=recover();err!=nil{
			switch err.(type){
			case runtime.Error:
				log.Println("runtime err:",err,"Check that the key is correct")
			default:
				log.Println("error:",err)
			}
		}
	}()
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		return []byte{},err
	}
	plainText,err= rsa.DecryptPKCS1v15(rand.Reader, privateKey, cryptText)
	if err!=nil{
		return []byte{},err
	}
	return plainText,nil
}
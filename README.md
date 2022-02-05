# cred
常用加密：golang常用rsa/bcrypt/md5封装
# demo
```go
rsa:=NewRSA()
//生成RSA密钥对
err := rsa.GenerateKey(1024)

//RSA加密
text:=[]byte("hello world")
encryptText, err :=rsa.Encrypt(text,[]byte(publicKEY))
if err != nil {
  return 
}
//RSA解密
decryptText, _ :=rsa.Decrypt(encryptText,[]byte(privateKEY))
fmt.Printf("encryptText:%s\ndecryptText:%s\n",encryptText,decryptText)
```

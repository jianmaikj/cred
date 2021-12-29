package cred

import (
	"fmt"
	"testing"
)

const (
	privateKEY = `-----BEGIN PRIVATE KEY-----
MIICWwIBAAKBgQDCwhnuMFcvAo1sFGSw3/brBH1RoZYUDi23rAvW8XuR3zV7lcDv
yD/zUS/JLONEufVk9xXL6XBe79I9maZzch4m8YTLiEuedMCGg7NCX/w0y72TIRbe
ru2A+9BGi2bXAPf0kfs9SgTUzJnFOUEgoTRPMiTe69dv/pVEk/9vsXqlIQIDAQAB
AoGAEtNubD7R/qJRh1QFU6+ij580efKTrkaclreTwWhAWJ5p7hNQvhzuSZfcdabY
JSmQlfzoZ3ieOx8VhIzYsSojzdZ8PCAibo6ENChromkEBTPfvItcOGQV5S471UxV
epva2nuBgdE1usbea96r69uK79ZByXifEskM5yFJLzLm7G0CQQDI8VEd1l4Lcap/
y2z/j4k5wrl/+tb+cI9AR6UkMW19A+LWEJzIFcsjG3oRUhhWvBcooTEcl4sgrVvH
tC2us7a7AkEA+B79nRHavad+HduD4gsxAnwlhCUNNplC3xgllDHHDVSD61vYCvvZ
L5EtAY1hIni6mGu5/zwzQg42DhwepfML0wJAJxF0Gf8BzevNoIcmrBF2Nrq7Yjyc
7F9qGq9TVlQd236T7dbRBL3n/u1qsE9r2AE953JfoBV+wV2zSIKplxvt3QJAXEd6
X05ioiKG/yfDFvC+m4P8l/cHwQSjIEHkyAbzVvvVzFTuNreQAObnbCpes+lezI1U
MZ78GVcbt4exxlpsTQJAQPNB6MBpawIcAyWHs0MqV+92go14l1mqgyHgLy3JFsxL
1dLkupSbaiOlT0JZvWoSrBhreY5quoEfhqzqHZWKsw==
-----END PRIVATE KEY-----`
	publicKEY = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCwhnuMFcvAo1sFGSw3/brBH1R
oZYUDi23rAvW8XuR3zV7lcDvyD/zUS/JLONEufVk9xXL6XBe79I9maZzch4m8YTL
iEuedMCGg7NCX/w0y72TIRberu2A+9BGi2bXAPf0kfs9SgTUzJnFOUEgoTRPMiTe
69dv/pVEk/9vsXqlIQIDAQAB
-----END PUBLIC KEY-----`
)
func Test(t *testing.T) {
	rsa:=NewRSA()
	//err := rsa.GenerateKey(1024)
	text:=[]byte("hello world")
	encryptText, err :=rsa.Encrypt(text,[]byte(publicKEY))
	if err != nil {
		return 
	}
	decryptText, _ :=rsa.Decrypt(encryptText,[]byte(privateKEY))
	fmt.Printf("encryptText:%s\ndecryptText:%s\n",encryptText,decryptText)
}

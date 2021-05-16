package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	myMD5 "github.com/lppgo/my_test/033_crypto/md5"
	myRSA "github.com/lppgo/my_test/033_crypto/rsa"
)

func main() {
	// 摘要算法
	MD5Example()

	// 对称加密

	// 非对称加密
	RSAExample()

}

// 1: RSA进行加解密.
// 2: RSA进行数字签名和签名校验.
func RSAExample() {
	msg := []byte("lucas的加密信息是...")
	// 生成密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("rsa生成密钥失败:", err.Error())
	}
	publicKey := privateKey.PublicKey

	// RSA加解密
	encryptedBytes := myRSA.RSAEncrypt(publicKey, msg)
	fmt.Println("encrypted bytes: ", encryptedBytes)
	decryptedBytes := myRSA.RSADecrypt(*privateKey, encryptedBytes)
	fmt.Println("decrypted message:", string(decryptedBytes))

	// RSA进行数字签名sign
	// 注意，只有拥有私钥的人才能对信息进行签名，但是有公钥的人可以验证它
	msg = []byte("lucas sign message")
	// 生成hash
	hash := sha256.New()
	_, err = hash.Write(msg)
	if err != nil {
		fmt.Println("hash.Write失败:", err.Error())
	}
	msgHashSum := hash.Sum(nil)
	//
	sign := myRSA.GenerateSign(*privateKey, msgHashSum)
	fmt.Println("生成的签名sign:", sign)
	ok := myRSA.CheckSign(publicKey, msgHashSum, sign)
	fmt.Println("对签名sign校验:", ok)
}

func MD5Example() {
	str := "123456"
	string1 := myMD5.GetMd5String1(str)
	fmt.Println(string1)
	string2 := myMD5.GetMd5String2([]byte(str))
	fmt.Println(string2)
}

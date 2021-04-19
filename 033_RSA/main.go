package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// 1: RSA进行加解密.
// 2: RSA进行数字签名和签名校验.
func main() {
	msg := []byte("lucas的加密信息是...")

	// 生成密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("rsa生成密钥失败:", err.Error())
	}
	publicKey := privateKey.PublicKey

	// RSA加解密
	encryptedBytes := RSAEncrypt(publicKey, msg)
	fmt.Println("encrypted bytes: ", encryptedBytes)
	decryptedBytes := RSADecrypt(*privateKey, encryptedBytes)
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
	sign := generateSign(*privateKey, msgHashSum)
	fmt.Println("生成的签名sign:", sign)
	ok := checkSign(publicKey, msgHashSum, sign)
	fmt.Println("对签名sign校验:", ok)

}

// RSAEncrypt .
func RSAEncrypt(publicKey rsa.PublicKey, msg []byte) []byte {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, msg, nil)
	if err != nil {
		fmt.Println("rsa.EncryptOAEP 加密失败:", err.Error())
	}
	return encryptedBytes
}

// RSADecrypt .
func RSADecrypt(privateKey rsa.PrivateKey, encryptedBytes []byte) []byte {
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		fmt.Println("privateKey.Decrypt 解密失败:", err.Error())
	}
	return decryptedBytes
}

// 生成签名 .
func generateSign(privateKey rsa.PrivateKey, msgHashSum []byte) []byte {
	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message
	signature, err := rsa.SignPSS(rand.Reader, &privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		fmt.Println("rsa.SignPSS生成Sign失败:", err.Error())
	}
	return signature
}

// 校验签名 .
func checkSign(publicKey rsa.PublicKey, msgHashSum, signature []byte) bool {
	err := rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("sing签名校验失败: ", err.Error())
		return false
	}
	return true
}

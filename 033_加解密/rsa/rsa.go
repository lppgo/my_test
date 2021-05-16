package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

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
func GenerateSign(privateKey rsa.PrivateKey, msgHashSum []byte) []byte {
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
func CheckSign(publicKey rsa.PublicKey, msgHashSum, signature []byte) bool {
	err := rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("sing签名校验失败: ", err.Error())
		return false
	}
	return true
}

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	"github.com/lppgo/my_test/033_crypto/myaes"
	"github.com/lppgo/my_test/033_crypto/mymd5"
	"github.com/lppgo/my_test/033_crypto/myrsa"
)

func main() {
	// 摘要算法
	// MD5Example()

	// 对称加密
	AESExample()

	// 非对称加密
	// RSAExample()

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
	encryptedBytes := myrsa.RSAEncrypt(publicKey, msg)
	fmt.Println("encrypted bytes: ", encryptedBytes)
	decryptedBytes := myrsa.RSADecrypt(*privateKey, encryptedBytes)
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
	sign := myrsa.GenerateSign(*privateKey, msgHashSum)
	fmt.Println("生成的签名sign:", sign)
	ok := myrsa.CheckSign(publicKey, msgHashSum, sign)
	fmt.Println("对签名sign校验:", ok)
}

// md5摘要算法示例
func MD5Example() {
	str := "123456"
	string1 := mymd5.GetMd5String1(str)
	fmt.Println(string1)
	string2 := mymd5.GetMd5String2([]byte(str))
	fmt.Println(string2)
}

// AES 加解密
func AESExample() {
	plaintext := []byte("Lucas is currently the best DisneyPlus show")
	key := []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")
	//加密
	ciphertext, _ := myaes.AesEncrypt(plaintext, key)
	//解密
	plaintext2, _ := myaes.AesDecrypt(ciphertext, key)

	fmt.Printf("plaintext:%v\n", string(plaintext2))

	// ciphertextStr, _ := myaes.EncryptByAes(ciphertext)
	// plaintextStr, _ := myaes.EncryptByAes(plaintext2)
	// fmt.Printf("ciphertext:%v\n", ciphertextStr)
}

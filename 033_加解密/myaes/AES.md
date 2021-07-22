[toc]

# 1：AES 概述

密码学中的高级加密标准（Advanced Encryption Standard，AES），又称 Rijndael 加密法，是美国联邦政府采用的一种区块加密标准。这个标准用来替代原先的 DES（Data Encryption Standard），已经被多方分析且广为全世界所使用。AES 中常见的有三种解决方案，分别为 AES-128、AES-192 和 AES-256。如果采用真正的 128 位加密技术甚至 256 位加密技术，蛮力攻击要取得成功需要耗费相当长的时间。

# 2：AES 有五种加密模式

- **ECB:** 电码本模式（Electronic Codebook Book (ECB)）、
- **CCB:** 密码分组链接模式（Cipher Block Chaining (CBC)）、
- **CTR:** 计算器模式（Counter (CTR)）、
- **CFB:** 密码反馈模式（Cipher FeedBack (CFB)）
- **OFB:** 输出反馈模式（Output FeedBack (OFB)）

# 3：AES 加密代码

## 3.1 ECB 模式

```go
package main

import (
    "crypto/aes"
    "fmt"
)

func AESEncrypt(src []byte, key []byte) (encrypted []byte) {
    cipher, _ := aes.NewCipher(generateKey(key))
    length := (len(src) + aes.BlockSize) / aes.BlockSize
    plain := make([]byte, length*aes.BlockSize)
    copy(plain, src)
    pad := byte(len(plain) - len(src))
    for i := len(src); i < len(plain); i++ {
        plain[i] = pad
    }
    encrypted = make([]byte, len(plain))
    // 分组分块加密
    for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
        cipher.Encrypt(encrypted[bs:be], plain[bs:be])
    }

    return encrypted
}

func AESDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
    cipher, _ := aes.NewCipher(generateKey(key))
    decrypted = make([]byte, len(encrypted))
    //
    for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
        cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
    }

    trim := 0
    if len(decrypted) > 0 {
        trim = len(decrypted) - int(decrypted[len(decrypted)-1])
    }

    return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
    genKey = make([]byte, 16)
    copy(genKey, key)
    for i := 16; i < len(key); {
        for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
            genKey[j] ^= key[i]
        }
    }
    return genKey
}
func main()  {

    source:="hello world"
    fmt.Println("原字符：",source)
    //16byte密钥
    key:="1443flfsaWfdas"
    encryptCode:=AESEncrypt([]byte(source),[]byte(key))
    fmt.Println("密文：",string(encryptCode))

    decryptCode:=AESDecrypt(encryptCode,[]byte(key))

    fmt.Println("解密：",string(decryptCode))
}
```

## 3.2 CBC 模式

```go
package main
import(
    "bytes"
    "crypto/aes"
    "fmt"
    "crypto/cipher"
    "encoding/base64"
)
func main() {
    orig := "hello world"
    key := "0123456789012345"
    fmt.Println("原文：", orig)
    encryptCode := AesEncrypt(orig, key)
    fmt.Println("密文：" , encryptCode)
    decryptCode := AesDecrypt(encryptCode, key)
    fmt.Println("解密结果：", decryptCode)
}
func AesEncrypt(orig string, key string) string {
    // 转成字节数组
    origData := []byte(orig)
    k := []byte(key)
    // 分组秘钥
    // NewCipher该函数限制了输入k的长度必须为16, 24或者32
    block, _ := aes.NewCipher(k)
    // 获取秘钥块的长度
    blockSize := block.BlockSize()
    // 补全码
    origData = PKCS7Padding(origData, blockSize)
    // 加密模式
    blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
    // 创建数组
    cryted := make([]byte, len(origData))
    // 加密
    blockMode.CryptBlocks(cryted, origData)
    return base64.StdEncoding.EncodeToString(cryted)
}
func AesDecrypt(cryted string, key string) string {
    // 转成字节数组
    crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
    k := []byte(key)
    // 分组秘钥
    block, _ := aes.NewCipher(k)
    // 获取秘钥块的长度
    blockSize := block.BlockSize()
    // 加密模式
    blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
    // 创建数组
    orig := make([]byte, len(crytedByte))
    // 解密
    blockMode.CryptBlocks(orig, crytedByte)
    // 去补全码
    orig = PKCS7UnPadding(orig)
    return string(orig)
}
//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
    padding := blocksize - len(ciphertext)%blocksize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}
//去码
func PKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}
```

## 3.3 CRT 模式

```go
package main

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "fmt"
)
//加密
func aesCtrCrypt(plainText []byte, key []byte) ([]byte, error) {

    //1. 创建cipher.Block接口
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    //2. 创建分组模式，在crypto/cipher包中
    iv := bytes.Repeat([]byte("1"), block.BlockSize())
    stream := cipher.NewCTR(block, iv)
    //3. 加密
    dst := make([]byte, len(plainText))
    stream.XORKeyStream(dst, plainText)

    return dst, nil
}


func main() {
    source:="hello world"
    fmt.Println("原字符：",source)

    key:="1443flfsaWfdasds"
    encryptCode,_:=aesCtrCrypt([]byte(source),[]byte(key))
    fmt.Println("密文：",string(encryptCode))

    decryptCode,_:=aesCtrCrypt(encryptCode,[]byte(key))

    fmt.Println("解密：",string(decryptCode))
}
```

## 3.4 CFB 模式

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
)
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
    block, err := aes.NewCipher(key)
    if err != nil {
        //panic(err)
    }
    encrypted = make([]byte, aes.BlockSize+len(origData))
    iv := encrypted[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        //panic(err)
    }
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
    return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
    block, _ := aes.NewCipher(key)
    if len(encrypted) < aes.BlockSize {
        panic("ciphertext too short")
    }
    iv := encrypted[:aes.BlockSize]
    encrypted = encrypted[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(encrypted, encrypted)
    return encrypted
}
func main() {
    source:="hello world"
    fmt.Println("原字符：",source)
    key:="ABCDEFGHIJKLMNO1"//16位
    encryptCode:=AesEncryptCFB([]byte(source),[]byte(key))
    fmt.Println("密文：",hex.EncodeToString(encryptCode))
    decryptCode:=AesDecryptCFB(encryptCode,[]byte(key))

    fmt.Println("解密：",string(decryptCode))
}
```

## 3.5 OFB 模式

```go
package main

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
)
func aesEncryptOFB( data[]byte,key []byte) ([]byte, error) {
    data = PKCS7Padding(data, aes.BlockSize)
    block, _ := aes.NewCipher([]byte(key))
    out := make([]byte, aes.BlockSize + len(data))
    iv := out[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewOFB(block, iv)
    stream.XORKeyStream(out[aes.BlockSize:], data)
    return out, nil
}

func aesDecryptOFB( data[]byte,key []byte) ([]byte, error) {
    block, _ := aes.NewCipher([]byte(key))
    iv  := data[:aes.BlockSize]
    data = data[aes.BlockSize:]
    if len(data) % aes.BlockSize != 0 {
        return nil, fmt.Errorf("data is not a multiple of the block size")
    }

    out := make([]byte, len(data))
    mode := cipher.NewOFB(block, iv)
    mode.XORKeyStream(out, data)

    out= PKCS7UnPadding(out)
    return out, nil
}
//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
    padding := blocksize - len(ciphertext)%blocksize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}
//去码
func PKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}
func main() {
    source:="hello world"
    fmt.Println("原字符：",source)
    key:="1111111111111111"//16位  32位均可
    encryptCode,_:=aesEncryptOFB([]byte(source),[]byte(key))
    fmt.Println("密文：",hex.EncodeToString(encryptCode))
    decryptCode,_:=aesDecryptOFB(encryptCode,[]byte(key))

    fmt.Println("解密：",string(decryptCode))
}
```

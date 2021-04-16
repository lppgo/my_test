- [一： 一些整理的函数](#一-一些整理的函数)
    - [1: string和[]byte高效互转](#1-string和byte高效互转)
    - [2: 结构体和[]byte互转](#2-结构体和byte互转)
    - [3: int和[]byte互转](#3-int和byte互转)
    - [4: go Generate UUID](#4-go-generate-uuid)
    - [5: IsEmpty()判断给的值是否为空](#5-isempty判断给的值是否为空)
    - [6: 二分法对Slice进行插入排序](#6-二分法对slice进行插入排序)
    - [7: RemoteIp 返回远程客户端的 IP，如 192.168.1.1   ; 将 IPv4 字符串形式转为 uint32](#7-remoteip-返回远程客户端的-ip如-19216811----将-ipv4-字符串形式转为-uint32)
    - [8：UTF-8和GBK转换](#8utf-8和gbk转换)
    - [9: RSA加密解密](#9rsa加密解密)
    - [10: MD5加密解密](#10md5加密解密)
* [四：其他](#四其他)

```go
1: string和[]byte高效互转
2: 结构体和[]byte互转
3: int和[]byte互转
4: go Generate UUID
5: IsEmpty()判断给的值是否为空
6: 二分法对Slice进行插入排序
7: RSA加密解密
8：获取远程客户端的IP,讲IPV4转uint32
9: UTF-8和GBK转换
```

# 一： 一些整理的函数

### 1: string和[]byte高效互转
```go
// string高效转换为[]byte
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// []byte高效转换为string
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
```
### 2: 结构体和[]byte互转
```go
type MyStruct struct {
	A int
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}
```
### 3: int和[]byte互转
```go
func Int2Bytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}
func Bytes2Int(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}
```
### 4: go Generate UUID
```go
//"github.com/google/uuid"
func GenerateUUID() {
	id := uuid.New()
	fmt.Printf("%s %s\n", id, id.Version().String())
	id2, _ := uuid.NewRandom()
	fmt.Printf("%s %s\n", id2, id2.Version().String())
}
```
### 5: IsEmpty()判断给的值是否为空
```go
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch value := value.(type) {
	case int:
		return value == 0
	case int8:
		return value == 0
	case int16:
		return value == 0
	case int32:
		return value == 0
	case int64:
		return value == 0
	case uint:
		return value == 0
	case uint8:
		return value == 0
	case uint16:
		return value == 0
	case uint32:
		return value == 0
	case uint64:
		return value == 0
	case float32:
		return value == 0
	case float64:
		return value == 0
	case bool:
		return value == false
	case string:
		return value == ""
	case []byte:
		return len(value) == 0
	default:
		// Finally using reflect.
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Array:
			return rv.Len() == 0

		case reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return true
			}
		}
	}
	return false
}
```
### 6: 二分法对Slice进行插入排序
```go
type muxEntry struct {
	h       http.Handler
	pattern string
}

func appendSorted(es []muxEntry, e muxEntry) []muxEntry {
	n := len(es)
	// 排序之后，进行二分搜索，若找见符合条件=true,返回index，如果不符合条件,i==n
	i := sort.Search(n, func(i int) bool {
		return len(es[i].pattern) < len(e.pattern) //sorted from longtest to shortest
		//return len(es[i].pattern) > len(e.pattern)
	})
	if i == n {
		return append(es, e)
	}
	// we now know that i points at where we want to insert
	es = append(es, muxEntry{}) // try to grow the slice in place, any entry works.
	copy(es[i+1:], es[i:])      // Move shorter entries down
	es[i] = e
	return es
}
```

### 7: RemoteIp 返回远程客户端的 IP，如 192.168.1.1   ; 将 IPv4 字符串形式转为 uint32
```go
// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}
//IP转换为uint32类型
func IPToUInt32(ipnr net.IP) uint32 {
    bits := strings.Split(ipnr.String(), ".")
 
    b0, _ := strconv.Atoi(bits[0])
    b1, _ := strconv.Atoi(bits[1])
    b2, _ := strconv.Atoi(bits[2])
    b3, _ := strconv.Atoi(bits[3])
 
    var sum uint32
 
    sum += uint32(b0) << 24
    sum += uint32(b1) << 16
    sum += uint32(b2) << 8
    sum += uint32(b3)
 
    return sum
}

//uint32类型转换为IP
func UInt32ToIP(intIP uint32) net.IP {
    var bytes [4]byte
    bytes[0] = byte(intIP & 0xFF)
    bytes[1] = byte((intIP >> 8) & 0xFF)
    bytes[2] = byte((intIP >> 16) & 0xFF)
    bytes[3] = byte((intIP >> 24) & 0xFF)
    return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
```

### 8：UTF-8和GBK转换
```go
package main

import (
    "bytes"
    "fmt"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
)

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

```
### 9:RSA加密解密

```go
//加密
func RAS_Encrypt() {
	//生成密钥
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey = privateKey.PublicKey
	//加密
	encryptedBytes, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, msg, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("encrypted bytes: ", encryptedBytes)
}
// 解密
func RAS_Decrypt() {
	var err error
	decryptedBytes, err = privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypted message:", string(decryptedBytes))
}
```
### 10:MD5加密解密
```go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

//方式一
func getMd5String1(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

//方式二
func getMd5String2(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func main() {
	str := "123456"
	string1 := getMd5String1(str)
	fmt.Println(string1)
	string2 := getMd5String2([]byte(str))
	fmt.Println(string2)
}
```
# 四：其他


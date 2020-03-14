package util

import (
	"unsafe"
)

//https://www.cnblogs.com/shuiyuejiangnan/p/9707066.html

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

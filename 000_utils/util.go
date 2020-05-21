package util

import (
	"unsafe"
	"reflect"
)

//https://www.cnblogs.com/shuiyuejiangnan/p/9707066.html
//--------------------------string和[]byte高效互转--------------------------------------
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

//-----------------------------结构体和[]byte互转----------------------------------------
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

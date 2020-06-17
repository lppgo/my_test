package util

import (
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func init() {
	//1: string和[]byte高效互转
	//2: 结构体和[]byte互转
	//3: IsEmpty()判断给的值是否为空
	//4: 二分法对Slice进行插入排序
	//5: 生成 len=18的SessionID
	//6:
	//7:
	//8:
	//9:
	//10:
	//11:
	//12:
}

//--------------------------string和[]byte高效互转------------------------------------------1-------------------------------------
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

//-----------------------------结构体和[]byte互转-------------------------------------------2------------------------------------
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

// -----------------------------IsEmpty()判断给的值是否为空----------------------------------------3--------------------------------
// IsEmpty checks whether given <value> empty.
// It returns true if <value> is in: 0, nil, false, "", len(slice/map/chan) == 0.Or else it returns true.
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

// -----------------------------二分法对Slice进行插入排序----------------------------------------4--------------------------------
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

// -----------------------------生成 len=18的SessionID----------------------------------------5--------------------------------
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 生成 len=18的SessionID
func GenerateSessionID() string {
	uninxNano := time.Now().UnixNano()
	return strings.ToUpper(strconv.FormatInt(uninxNano, 36) + Str(6))
}

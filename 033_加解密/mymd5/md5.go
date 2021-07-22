package mymd5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

//方式一
func getMd5String1(str string) string {
	h := md5.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		log.Fatal(err)
	}
	// arr := h.Sum(nil)
	// return fmt.Sprintf("%x", arr)
	arr := hex.EncodeToString(h.Sum(nil))
	return arr
}

//方式二
func GetMd5String2(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

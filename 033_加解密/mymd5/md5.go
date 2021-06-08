package mymd5

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

//方式一
func GetMd5String1(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

//方式二
func GetMd5String2(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

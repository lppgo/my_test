package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// destination
//
func main() {
	log.Println("-------------- main start ... -----------")

	src := "./test.txt"
	dst := "./test222.txt"
	//n, err := copy1(src, dst)
	//err := copy2(src, dst)
	err := copy3(src, dst)

	fmt.Println(err)

	log.Println("---------------- main end ... -----------")

}

// copy 方法1 : 方法将使用标准Go库的 io.Copy()函数。以下是使用io.Copy()实现的拷贝文件
func copy1(src, dst string) (int64, error) {
	source, err := os.Open(src)
	if err != nil {
		err = fmt.Errorf("OpenErr:", err.Error())
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		err = fmt.Errorf("CreateErr:", err.Error())
		return 0, err
	}
	defer destination.Close()
	//
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// 方法2 ： 使用ioutil包中的 ioutil.WriteFile()和 ioutil.ReadFile()，但由于使用一次性读取文件，
//再一次性写入文件的方式，所以该方法不适用于大文件，容易内存溢出
func copy2(src, dst string) (err error) {
	fBytes, err := ioutil.ReadFile(src)
	if err != nil {
		err = fmt.Errorf("ReadFileErr:%s", err.Error())
		return
	}

	err = ioutil.WriteFile(dst, fBytes, 0644)
	if err != nil {
		err = fmt.Errorf("WriteFileErr:%s", err.Error())
	}
	return err
}

// 方法3 : 使用os包的os.Read()和os.Write(),此方法是按块读取文件，块的大小也会影响程序性能
func copy3(src, dst string) (err error) {
	BufferSize := 1024
	buf := make([]byte, BufferSize)

	source, err := os.Open(src)
	if err != nil {
		err = fmt.Errorf("OpenErr:", err.Error())
		return err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		err = fmt.Errorf("CreateErr:", err.Error())
		return err
	}
	defer destination.Close()

	//
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err = destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return
}

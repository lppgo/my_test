package main

import (
	"fmt"
	"log"
	"time"
)

/*
	1: 不带参的闭包函数
	2: 带参的闭包函数
	3: 闭包函数内部修改外部的变量的话，外部变量的值会改变。闭包函数里面引用的是变量的地址
*/

func main() {
	log.Println("===start===")
	a := 10
	str := "milk"

	defer func() {
		a = 50
		fmt.Println("a 1 is:", a, "  str 1 is:", str)
		fmt.Println()
	}()

	defer func(aa int, ss string) {
		fmt.Println("a 2 is:", aa, "   str 2 is:", ss)
		fmt.Println()
	}(a, str)

	a = 100
	str = "interchange交换"

	time.Sleep(time.Second * 3)
	log.Println("===end===")
}

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
		fmt.Printf("a 1 is:%d ,  a 1 Addr:%p,  str 1 is:%s\n", a, &a, str)
	}()

	defer func(a int, str string) {
		fmt.Printf("a 2 is:%d ,  a 2 Addr:%p,  str 2 is:%s\n", a, &a, str)
	}(a, str)

	a = 100
	str = "interchange交换"

	fmt.Printf("a is:%d ,  a Addr:%p,  str is:%s\n", a, &a, str)

	time.Sleep(time.Second * 3)
	log.Println("===end===")
}

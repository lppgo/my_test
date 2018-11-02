package main

import (
	"log"
)

/* 回调函数：函数有一个参数是函数类型，这个函数就是回调函数 */

// 定义一个函数类型
type FuncType func(int, int) int

// 定义一个函数
func add(a, b int) int {
	return a + b
}

// 定义一个回调函数
func CallBack(a, b int, testFunc FuncType) (result int) {
	value := testFunc(a, b)
	return value
}

func main() {
	log.Println("----使用callback回调函数----")
	result := CallBack(4, 5, add)
	log.Println(result)
}

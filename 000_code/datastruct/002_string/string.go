package string

import "fmt"

func searchSubString() {
	str := "goodgoogle"
	subStr := "google"

	fmt.Println("普通方式字符串匹配 1:", searchSubStr(str, subStr))

	fmt.Println("KMP 字符串匹配    2:", searchSubStrKMP(str, subStr))

	s := "abcac"
	fmt.Println("计算模式字符串前缀表:", getNext(s))
}

// 匹配字符串

// 1：暴力法
func searchSubStr(str, subStr string) int {
	n, m := len(str), len(subStr)
	index := -1

	switch {
	case m == 0:
		index = 0
		return index
	case m > n:
		return index
	case m == n && str == subStr:
		index = 0
		return index
	case m == n && str != subStr:
		return index
	}
	//
	for i := 0; i <= n-m; i++ {
		if str[i:i+m] == subStr {
			index = i
			break
		}
		index = -1
	}
	return index
}

// 2：KMP算法
func searchSubStrKMP(str, subStr string) int {
	n, m := len(str), len(subStr)
	index := -1

	switch {
	case m == 0:
		index = 0
		return index
	case m > n:
		return index
	case m == n && str == subStr:
		index = 0
		return index
	case m == n && str != subStr:
		return index
	}

	// 计算模式串的前缀表
	next := getNext(subStr)
	i, j := 0, 0
	for i < len(str) && j < len(subStr) {
		if j == -1 || str[i] == subStr[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(subStr) {
		return i - j
	}
	return index
}

// 计算模式串的前缀表
// 模式串的前缀
// 模式串的后缀
// 最长相等的前后缀
func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	i, j := 0, -1
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

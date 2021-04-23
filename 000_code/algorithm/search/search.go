package search

import "fmt"

// search  搜索算法

func Search() {

	// 滑动窗口法 .
	// s := "asbiabihcas19io2"
	// t := "abc"
	// slidingWindows(s)
	// slidingWindows2(s, t)

	// 	二分搜索法.
	arr := []int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9}
	index1 := binarySearch1(arr, 8)
	fmt.Println("non-recursion binarySearch:", index1)

}

// 滑动窗口 .
// 使用滑动窗口法获取无重复最长子串
func slidingWindows(s string) {
	windows := make(map[byte]int, len(s))
	left, right := 0, 0
	res := 0 //记录不重复子串最大长度

	for right < len(s) { // right 右滑动 .
		rv := s[right]
		windows[rv]++
		right++

		for windows[rv] > 1 { //left 右滑动.
			lv := s[left]
			windows[lv]--
			left++
		}
		// 判断区间是否满足条件
		if right-left > res {
			res = right - left
		}
	}
	fmt.Printf("字符串:%s的最大不重复子串是:%s,长度是:%d\n", s, s[left:right], res)
}

// 最小覆盖子串: 返回s中包含t中元素的最小子串
func slidingWindows2(s, t string) {
	lens := len(s)
	lent := len(t)
	windows := make(map[byte]int, lens)
	target := make(map[byte]int, lent)
	for _, v := range t {
		target[byte(v)] = 1
	}
	validCount := 0 // 当 validCount == lent ,也就是target是windows的子集的时候，满足条件，开始右移
	left, right := 0, 0
	min := lens // 符合条件最小子串长度
	start := -1 // 默认起始位置

	for right < lens {
		rVal := s[right]
		windows[rVal]++
		right++
		if windows[rVal] == target[rVal] {
			validCount++
		}

		for validCount == lent { //
			lVal := s[left]
			if right-left <= min {
				min = right - left
				start = left
			}
			if windows[lVal] == target[lVal] {
				validCount--
			}
			windows[lVal]--
			left++
		}

	}
	fmt.Println(min)
	if start == -1 {
		fmt.Printf(" 字符串s:%s不包含子串t:%s元素", s, t)
	} else {
		fmt.Printf(" 字符串s:%s包含字符串t:%s元素的最小字串是:%s\n", s, t, s[start:start+min])
	}
}

// non-recursion .
func binarySearch1(arr []int, target int) int {
	noexist := -1
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) >> 1
		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] > target:
			right = mid
		case arr[mid] < target:
			left = mid + 1
		}
	}
	return noexist
}

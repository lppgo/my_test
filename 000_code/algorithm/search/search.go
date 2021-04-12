package search

import "fmt"

// search  搜索算法

func Search() {

	s := "asbiabihcas19io2"
	t := "abc"
	// slidingWindows(s)
	slidingWindows2(s, t)
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
	windows := make(map[byte]int, len(s))
	target := make(map[byte]int, len(t)) //
	for _, v := range t {
		target[byte(v)]++
	}

	left, right, valid := 0, 0, 0
	start := -1 // s
	res := len(s)

	for right < len(s) { // right右滑
		rval := s[right]
		windows[rval]++
		right++
		if windows[rval] == target[rval] {
			valid++
		}

		for valid == len(target) {
			lval := s[left]
			if right-left <= res { //是否更新结果值
				start = left
				res = right - left
			}
			if windows[lval] == target[lval] {
				valid--
			}
			windows[lval]--
			left++
		}
	}

	if start == -1 {
		fmt.Printf(" 字符串s:%s不包含子串t:%s元素", s, t)
	} else {
		fmt.Printf(" 字符串s:%s包含字符串t:%s元素的最小字串是:%s\n", s, t, s[start:start+res])
	}

}

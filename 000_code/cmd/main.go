package main

import (
	sort "github.com/lppgo/my_test/000_code/algorithm/sort"
	heap "github.com/lppgo/my_test/000_code/datastruct/008_heap"
)

func main() {
	//
	option := 2

	switch option {
	case 1:
		DataStructApp()
	case 2:
		ALgorithmApp()
	}
}

// 数据结构 App.
func DataStructApp() {
	heap.ExampleIntHeap()
}

// 算法 App.
func ALgorithmApp() {
	sort.Sort()
}

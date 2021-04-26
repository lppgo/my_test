package main

import (
	// mySearch "github.com/lppgo/my_test/000_code/algorithm/search"
	// mySort "github.com/lppgo/my_test/000_code/algorithm/sort"
	// myHeap "github.com/lppgo/my_test/000_code/datastruct/008_heap"

	"log"

	mySort "github.com/lppgo/my_test/000_code/algorithm/sort"
	mySkiplist "github.com/lppgo/my_test/000_code/datastruct/009_skiplist"
)

func main() {
	option := 1

	switch option {
	case 1:
		DataStructApp()
	case 2:
		AlgorithmApp()
	}

	log.Println("exit ...")
}

// 数据结构 App.
func DataStructApp() {
	// myHeap.ExampleIntHeap()

	mySkiplist.ExampleSkiplist()
}

// 算法 App.
func AlgorithmApp() {
	mySort.Sort()

	// mySearch.Search()
}

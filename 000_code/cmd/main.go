package main

import (
	// mySearch "github.com/lppgo/my_test/000_code/algorithm/search"
	// mySort "github.com/lppgo/my_test/000_code/algorithm/sort"
	// myHeap "github.com/lppgo/my_test/000_code/datastruct/008_heap"

	"log"

	mySort "github.com/lppgo/my_test/000_code/algorithm/sort"
	myHeap "github.com/lppgo/my_test/000_code/datastruct/008_heap"
)

func main() {
	option := 2

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
	myHeap.ExampleIntHeap()
}

// 算法 App.
func AlgorithmApp() {
	mySort.Sort()

	// mySearch.Search()
}

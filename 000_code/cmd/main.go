package main

import (
	// mySearch "github.com/lppgo/my_test/000_code/algorithm/search"
	// mySort "github.com/lppgo/my_test/000_code/algorithm/sort"
	// myHeap "github.com/lppgo/my_test/000_code/datastruct/008_heap"

	// myArray "github.com/lppgo/my_test/000_code/datastruct/001_array"
	// myTree "github.com/lppgo/my_test/000_code/datastruct/007_tree"

	"fmt"
	"log"

	mySort "github.com/lppgo/my_test/000_code/algorithm/sort"

	myTree "github.com/lppgo/my_test/000_code/datastruct/007_tree"
)

func main() {
	option := 1

	switch option {
	case 1:
		DataStructApp()
	case 2:
		AlgorithmApp()
	}

	fmt.Println()
	log.Println("exit ...")
}

// 数据结构 App.
func DataStructApp() {
	// myArray.ExampleArray()

	myTree.ExampleTree()

	// myHeap.ExampleIntHeap()

	// mySkiplist.ExampleSkiplist()
}

// 算法 App.
func AlgorithmApp() {
	mySort.Sort()

	// mySearch.Search()
}

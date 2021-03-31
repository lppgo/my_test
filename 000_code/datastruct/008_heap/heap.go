package heap

import (
	"container/heap"
	"fmt"
)

// 定义一个堆结构 min-heap 小根堆.
type IntHeap []int

// 实现 heap.interface 接口.
func (ih *IntHeap) Len() int {
	return len(*ih)
}

func (ih *IntHeap) Less(i, j int) bool {
	return (*ih)[i] < (*ih)[j]
}

func (ih *IntHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
}

// add x as element Len().
func (ih *IntHeap) Push(h interface{}) {
	*ih = append(*ih, h.(int))
}

// remove and return element Len() - 1.
func (ih *IntHeap) Pop() interface{} {
	n := len(*ih)
	x := (*ih)[n-1] // 返回删除的元素
	(*ih) = (*ih)[:n-1]
	return x
}

// 实现heap的操作.
func ExampleIntHeap() {
	h := &IntHeap{2, 1, 5}

	heap.Init(h)
	fmt.Printf("heap.Init() len:%d ,IntHeap:%v\n", len(*h), h)
	heap.Push(h, 3)
	heap.Push(h, 4)

	// 修改某个元素的value,heap自动调整
	i := 3
	heap.Fix(h, i)
	fmt.Printf("heap.Fix()  len:%d ,IntHeap:%v\n", len(*h), h)

	heap.Pop(h)
	fmt.Printf("heap.Pop()  len:%d ,IntHeap:%v\n", len(*h), h)

	heap.Remove(h, 3)
	fmt.Printf("heap.Remove() len:%d ,IntHeap:%v\n", len(*h), h)
}

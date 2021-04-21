package sort

import (
	"fmt"
)

func generateArrList() [][]int {
	return [][]int{
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 0
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 1
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 2
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 3
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 4
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 5
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 6
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 7
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 8
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 9
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 10
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 11
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 12
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 13
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 14
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 15
	}
}

// Sort 排序算法.
func Sort() {
	arr := generateArrList()

	// 选择排序: 时间复杂度O(n^2)，空间复杂度O(1),不稳定排序
	fmt.Println("SelectSort:", SelectSort(arr[0]))

	// 冒泡排序: 时间复杂度O(n^2)，空间复杂度O(1),稳定排序
	fmt.Println("BubbleSort:", BubbleSort(arr[1]))

	// 插入排序: 时间复杂度O(n^2)，空间复杂度O(1),稳定排序，适合整个数组大部分元素有序，个别元素无序的情况
	fmt.Println("InsertSort:", InsertSort(arr[2]))

	// 快速排序：时间复杂度O(nlogn),空间复杂度O(logn),不稳定排序
	// 快速排序1: 用基准值mid
	fmt.Println("QuickSort1:", QuickSort1(arr[3]))
	// 快速排序2: 用基准值mid + left, right指针
	fmt.Println("QuickSort2:", QuickSort2(arr[4]))

	// 归并排序：时间复杂度O(nlogn),空间复杂度O(n),稳定排序
	fmt.Println("MergeSort :", MergeSort(arr[5]))

	// 希尔排序：时间复杂度O(nlogn),空间复杂度O(1),不稳定排序
	fmt.Println("ShellSort :", ShellSort(arr[6]))

	// 堆排序：时间复杂度O(logn),空间复杂度O(1),不稳定排序
	fmt.Println("HeapSort  :", HeapSort(arr[7]))
}

//
func SelectSort(a []int) []int {
	min := 0
	for i := 0; i < len(a); i++ {
		min = a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]              //
				a[i], a[j] = a[j], a[i] //
			}
		}
	}
	return a
}

//
func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

//
func InsertSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		p := i
		for j := p - 1; j >= 0; j-- {
			if a[p] < a[j] {
				a[p], a[j] = a[j], a[p]
				p = j //
			}
		}
	}
	return a
}

//
func QuickSort1(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	p := a[len(a)>>1]

	lt := make([]int, 0, len(a))
	eq := make([]int, 0, len(a))
	gt := make([]int, 0, len(a))
	result := make([]int, 0, len(a))

	for i := 0; i < len(a); i++ {
		switch {
		case a[i] > p:
			gt = append(gt, a[i])
		case a[i] == p:
			eq = append(eq, a[i])
		case a[i] < p:
			lt = append(lt, a[i])
		}
	}
	// 递归 recursion
	lt = QuickSort1(lt)
	gt = QuickSort1(gt)

	result = append(result, lt...)
	result = append(result, eq...)
	result = append(result, gt...)

	return result
}

//
func QuickSort2(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	p := a[0]
	left, right := 0, len(a)-1

	for i := 1; i <= right; { // 注意边界条件
		if left >= right {
			break
		}

		if a[i] > p {
			a[i], a[right] = a[right], a[i]
			right--
		} else {
			a[i], a[left] = a[left], a[i]
			left++
			i++
		}
	}
	QuickSort2(a[:left])
	QuickSort2(a[left+1:])

	return a
}

//
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 获取分区位置
	p := len(nums) / 2
	// 通过递归分区
	left := MergeSort(nums[0:p])
	right := MergeSort(nums[p:])
	// 排序后合并
	return merge(left, right)
}

// 排序合并 .
func merge(left []int, right []int) []int {
	i, j := 0, 0
	// 用于存放结果集
	var result []int
	for {
		// 任何一个区间遍历完，则退出
		if i >= len(left) || j >= len(right) {
			break
		}
		// 对所有区间数据进行排序
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 如果左侧区间还没有遍历完，将剩余数据放到结果集
	if i != len(left) {
		result = append(result, left[i:]...)
	}

	// 如果右侧区间还没有遍历完，将剩余数据放到结果集
	if j != len(right) {
		result = append(result, right[j:]...)
	}

	// 返回排序后的结果集
	return result
}

//
func ShellSort(a []int) []int {
	// 寻找合适的步长间隔h ，3h+1或者n/2
	h := 1
	for h < len(a)/3 {
		h = h*3 + 1
	}

	for step := h; step > 0; step /= 3 { // 步长step缩小
		for i := step; i < len(a); i++ {
			for j := i - step; j >= 0 && a[j] > a[j+step]; j -= step {
				a[j], a[j+step] = a[j+step], a[j]
			}
		}
	}
	return a
}

//
func HeapSort(a []int) []int {
	// buildMaxHeap(a)
	end := len(a) - 1 // end 最后一个元素的index
	if end < 1 {
		return a
	}
	// p 父节点index
	for parent := end / 2; parent >= 0; parent-- {
		sink(a, parent, end)
	}
	fmt.Printf("maxheap build完成:%v\n", a)

	// pop
	result := make([]int, len(a))
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = pop(&a)
	}
	return result
}

func buildMaxHeap(a []int) {
	n := len(a) - 1 // 最后一个元素的index
	if n < 1 {
		return
	}
	// p 父节点index
	for p := (n - 1) / 2; p >= 0; p-- {
		sink(a, p, n)
	}
}

func pop(a *[]int) (root int) {
	n := len(*a) - 1
	if n > 0 {
		newheap := make([]int, n)
		root = (*a)[0]
		(*a)[0], (*a)[n] = (*a)[n], (*a)[0]
		copy(newheap, (*a)[:n])
		buildMaxHeap(newheap)
		*a = newheap
	}
	return
}

// 堆节点下沉 大的value下沉.
func sink(a []int, parent, end int) {
	for end > 0 {
		left := parent*2 + 1  // 左孩子
		right := parent*2 + 2 // 右孩子
		larger := left        // 存放较大的那一个孩子的index

		if left > end { // 防止索引溢出,判断是否存在left节点
			break
		}
		if right <= end && a[left] < a[right] {
			larger = right
		}

		if a[parent] >= a[larger] {
			break
		}
		a[parent], a[larger] = a[larger], a[parent]
		parent = larger
	}
}

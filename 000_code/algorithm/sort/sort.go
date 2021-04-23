package sort

import (
	"fmt"
	"strconv"
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
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 16
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 17
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 18
		{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}, // 19
	}
}

// Sort 排序算法.
func Sort() {
	arr := generateArrList()

	// 选择排序: 时间复杂度O(n^2)，空间复杂度O(1),不稳定排序
	fmt.Println("SelectSort   :", SelectSort(arr[0]))

	// 冒泡排序: 时间复杂度O(n^2)，空间复杂度O(1),稳定排序
	fmt.Println("BubbleSort   :", BubbleSort(arr[1]))

	// 插入排序: 时间复杂度O(n^2)，空间复杂度O(1),稳定排序，适合整个数组大部分元素有序，个别元素无序的情况
	fmt.Println("InsertSort   :", InsertSort(arr[2]))

	// 快速排序：时间复杂度O(nlogn),空间复杂度O(logn),不稳定排序
	// 快速排序1: 用基准值mid
	fmt.Println("QuickSort1   :", QuickSort1(arr[3]))
	// 快速排序2: 用基准值mid + left, right指针
	fmt.Println("QuickSort2   :", QuickSort2(arr[4]))

	// 归并排序：时间复杂度O(nlogn),空间复杂度O(n),稳定排序
	// 归并排序1: 使用递归
	fmt.Println("MergeSort1   :", MergeSort1(arr[5]))
	// 归并排序2：使用非递归
	// fmt.Println("MergeSort2 :", MergeSort2(arr[6]))

	// 希尔排序：时间复杂度O(nlogn),空间复杂度O(1),不稳定排序
	fmt.Println("ShellSort    :", ShellSort(arr[7]))

	// 堆排序：时间复杂度O(logn),空间复杂度O(1),不稳定排序
	fmt.Println("HeapSort     :", HeapSort(arr[8]))

	// 桶排序: 时间复杂度O(n+k),空间复杂度O(n+k),稳定排序，非原地排序
	// 数据有n个，我们把它们分在m个桶中，这样每个桶里的数据就是k = n / m
	// 桶排序的应用场景十分严格，数据分布比较均匀
	fmt.Println("BucketSort   :", BucketSort(arr[9]))

	// 计数排序：时间复杂度O(n+k),空间复杂度O(n+k),稳定排序，非原地排序
	fmt.Println("CountingSort :", CountingSort(arr[10]))

	// 基数排序：时间复杂度O(n+k),空间复杂度O(n+k),稳定排序，非原地排序
	// 基数排序(Radix sort) 是一种非比较型整数排序算法，其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较。
	// 由于整数也可以表达字符串（比如名字或日期）和特定格式的浮点数，所以基数排序也不是只能使用于整数。
	fmt.Println("RadixSort    :", RadixSort(arr[11]))

	// 拓扑排序
	/*
			注意:
				1) 只有有向无环图才存在拓扑序列;
				2) 对于一个 DAG (Directed Acyclic Graph), 可能存在多个拓扑序列;

		 拓扑序列算法思想
		(1)从有向图中选取一个没有前驱 (即入度为 0) 的顶点，并输出之;
		(2) 从有向图中删去此顶点以及所有以它为尾的弧;
		重复上述两步，直至图空，或者图不空但找不到无前驱的顶点为止。
	*/

	fmt.Println("TopoSort     :", Topo())

}

//
func SelectSort(arr []int) []int {
	min := 0
	for i := 0; i < len(arr); i++ {
		min = arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]                    //
				arr[i], arr[j] = arr[j], arr[i] //
			}
		}
	}
	return arr
}

//
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//
func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		p := i
		for j := p - 1; j >= 0; j-- {
			if arr[p] < arr[j] {
				arr[p], arr[j] = arr[j], arr[p]
				p = j //
			}
		}
	}
	return arr
}

//
func QuickSort1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	p := arr[len(arr)>>1]

	lt := make([]int, 0, len(arr))
	eq := make([]int, 0, len(arr))
	gt := make([]int, 0, len(arr))
	result := make([]int, 0, len(arr))

	for i := 0; i < len(arr); i++ {
		switch {
		case arr[i] > p:
			gt = append(gt, arr[i])
		case arr[i] == p:
			eq = append(eq, arr[i])
		case arr[i] < p:
			lt = append(lt, arr[i])
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
func QuickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	p := arr[0]
	left, right := 0, len(arr)-1

	for i := 1; i <= right; { // 注意边界条件
		if left >= right {
			break
		}

		if arr[i] > p {
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else {
			arr[i], arr[left] = arr[left], arr[i]
			left++
			i++
		}
	}
	QuickSort2(arr[:left])
	QuickSort2(arr[left+1:])

	return arr
}

//
func MergeSort1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 获取分区位置
	p := len(nums) / 2
	// 通过递归分区
	left := MergeSort1(nums[0:p]) // 先将左边排好
	right := MergeSort1(nums[p:]) // 再讲右边排好
	// 排序后合并
	return merge(left, right)
}

// 排序合并 .
func merge(left []int, right []int) []int {
	i, j := 0, 0
	var result []int // 用于存放结果集
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

// 归并排序2:非递归
// func MergeSort2(nums []int) []int {
// 	n := len(nums)
// 	res := make([]int, 0, n)
// 	fmt.Println(n)
// 	// 	子数组大小分别为1,2,4,8...
// 	// 	刚开始合并的数组大小是1,接着是2，接着是4...
// 	for i := 1; i < n; i += i {}
// 	return res
// }

//
func ShellSort(arr []int) []int {
	// 寻找合适的步长间隔h ，3h+1或者n/2
	h := 1
	for h < len(arr)/3 {
		h = h*3 + 1
	}

	for step := h; step > 0; step /= 3 { // 步长step缩小
		for i := step; i < len(arr); i++ {
			for j := i - step; j >= 0 && arr[j] > arr[j+step]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
	return arr
}

//
func HeapSort(arr []int) []int {
	// buildMaxHeap(arr)
	end := len(arr) - 1 // end 最后一个元素的index
	if end < 1 {
		return arr
	}
	// p 父节点index
	for parent := end / 2; parent >= 0; parent-- {
		sink(arr, parent, end)
	}
	fmt.Printf("maxheap build完成:%v\n", arr)

	// pop
	result := make([]int, len(arr))
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = pop(&arr)
	}
	return result
}

func buildMaxHeap(arr []int) {
	n := len(arr) - 1 // 最后一个元素的index
	if n < 1 {
		return
	}
	// p 父节点index
	for p := (n - 1) / 2; p >= 0; p-- {
		sink(arr, p, n)
	}
}

func pop(arr *[]int) (root int) {
	n := len(*arr) - 1
	if n > 0 {
		newheap := make([]int, n)
		root = (*arr)[0]
		(*arr)[0], (*arr)[n] = (*arr)[n], (*arr)[0]
		copy(newheap, (*arr)[:n])
		buildMaxHeap(newheap)
		*arr = newheap
	}
	return
}

// 堆节点下沉 大的value下沉.
func sink(arr []int, parent, end int) {
	for end > 0 {
		left := parent*2 + 1  // 左孩子
		right := parent*2 + 2 // 右孩子
		larger := left        // 存放较大的那一个孩子的index

		if left > end { // 防止索引溢出,判断是否存在left节点
			break
		}
		if right <= end && arr[left] < arr[right] {
			larger = right
		}

		if arr[parent] >= arr[larger] {
			break
		}
		arr[parent], arr[larger] = arr[larger], arr[parent]
		parent = larger
	}
}

func BucketSort(arr []int) []int {
	// 找出待排序数组的区间[min,max]，寻找数组的最大值与最小值
	n := len(arr)               // 桶数
	max := getMax(arr)          // 获取数组最大值
	buckets := make([][]int, n) // 二维切片
	// 分配入桶
	index := 0
	for i := 0; i < n; i++ {
		index = arr[i] * (n - 1) / max //assign bucket
		buckets[index] = append(buckets[index], arr[i])
	}
	// 桶内排序
	tmp := 0
	for i := 0; i < n; i++ {
		if len(buckets[i]) > 0 {
			sortInBucket(buckets[i])
			copy(arr[tmp:], buckets[i])
			tmp += len(buckets[i])
		}
	}
	return arr

}

// 获取数组最大值 .
func getMax(arr []int) (max int) {
	max = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return
}

// 桶内排序,用其他排序，比如快排都可以 .
func sortInBucket(arr []int) {
	QuickSort2(arr)
}

// 计数排序: 核心在于讲输入的数据转换为键值存储在额外开辟的数组空间
// 计数排序要求输入的数据必须是有确定范围的整数.
func CountingSort(arr []int) []int {
	max := getMax(arr)
	countArr := make([]int, max+1) //用来存放和计数
	for i := 0; i < len(arr); i++ {
		countArr[arr[i]]++
	}

	startIndex := 0
	for i := 0; i < len(countArr); i++ {
		for j := 1; j <= countArr[i]; j++ {
			arr[startIndex] = i
			startIndex++
		}
	}
	return arr
}

//
func RadixSort(arr []int) []int {
	max := getMax(arr)
	tmp := make([]int, len(arr), len(arr))
	count := new([10]int)
	radix := 1
	var i, j, k int
	for i = 0; i < max; i++ { // 进行max次排序
		for j = 0; j < 10; j++ {
			count[j] = 0
		}
		for j = 0; j < len(arr); j++ {
			k = (arr[j] / radix) % 10
			count[k]++
		}
		for j = 1; j < 10; j++ { //将tmp中的为准依次分配给每个桶
			count[j] = count[j-1] + count[j]
		}
		for j = len(arr) - 1; j >= 0; j-- {
			k = (arr[j] / radix) % 10
			tmp[count[k]-1] = arr[j]
			count[k]--
		}
		for j = 0; j < len(arr); j++ {
			arr[j] = tmp[j]
		}
		radix = radix * 10
	}
	return arr
}

// 问题描述：有一串数字1到5，按照下面的关于顺序的要求，重新排列并打印出来。
// 要求如下：2在5前出现，3在2前出现，4在1前出现，1在3前出现。
// 一般解决Topo排序的方案是采用DFS-深度优先算法

func Topo() []string {
	// edge 要求的拓扑顺序 4,1,3,2,5
	var edge = map[string]string{
		"2": "5",
		"3": "2",
		"4": "1",
		"1": "3",
	}

	result := make([]string, 0, 5)  // 结果数组
	visited := make([]string, 0, 5) // 已访问的数组

	for i := 1; i <= 5; i++ {
		TopoSort(edge, &result, &visited, strconv.Itoa(i))
	}
	reverse(result)
	return result
}

func TopoSort(edge map[string]string, result, visited *[]string, i string) {
	if !isVisited(*visited, i) {
		*visited = append(*visited, i)
		// fmt.Println("visited:", *visited, i)
		if _, ok := edge[i]; ok {
			TopoSort(edge, result, visited, edge[i])
		}
		*result = append(*result, i)
	}
}

// 检查是否存在已访问数组中
func isVisited(visited []string, element string) bool {
	isVisited := false
	for _, v := range visited {
		if v == element {
			isVisited = true
			break
		}
	}
	return isVisited
}
func reverse(result []string) {
	length := len(result)
	mid := length / 2
	for i := 0; i < mid; i++ {
		result[i], result[length-1-i] = result[length-i-1], result[i]
	}
}

package sort

import (
	"fmt"
)

// Sort 排序算法.
func Sort() {
	srcArray := []int{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}

	// 选择排序: 时间复杂度O(n^2)，空间复杂度O(1),不稳定排序
	SelectSort(srcArray)

	// 冒泡排序: 时间复杂度O(n^2)，空间复杂度O(1),稳定排序
	BubbleSort(srcArray)

	// 插入排序: 时间复杂度O(n^2)，空间复杂度O(1),稳定排序，适合整个数组大部分元素有序，个别元素无序的情况
	InsertSort(srcArray)

	// 快速排序：时间复杂度O(nlogn),空间复杂度O(logn),不稳定排序
	// 快速排序1: 用基准值mid
	fmt.Println("QuickSort1:", QuickSort1(srcArray))
	// 快速排序2: 用基准值mid + left, right指针
	fmt.Println("QuickSort2:", QuickSort2(srcArray))

}

//
func SelectSort(a []int) {
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
	fmt.Println("SelectSort:", a)
}

//
func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("BubbleSort:", a)
}

//
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i - 1; j >= 0; j-- {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println("InsertSort:", a)
}

//
func QuickSort1(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	mid := a[len(a)>>1]

	lt := make([]int, 0, len(a))
	eq := make([]int, 0, len(a))
	gt := make([]int, 0, len(a))
	result := make([]int, 0, len(a))

	for i := 0; i < len(a); i++ {
		switch {
		case a[i] > mid:
			gt = append(gt, a[i])
		case a[i] == mid:
			eq = append(eq, a[i])
		case a[i] < mid:
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

	mid := a[0]
	left, right := 0, len(a)-1

	for i := 1; i <= right; i++ {
		if left >= right {
			break
		}

		if a[i] > mid {
			a[i], a[right] = a[right], a[i]
			right--
		} else {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	QuickSort2(a[:left])
	QuickSort2(a[left+1:])

	return a
}

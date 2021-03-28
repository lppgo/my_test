package sort

import "fmt"

// Sort 排序算法
func Sort() {
	srcArray := []int{8, 5, 2, 6, 9, 6, 3, 1, 4, 0, 7}

	// 选择排序: 时间复杂度O(n^2)，空间复杂度O(1),不稳定排序
	SelectionSort(srcArray)

	// 冒泡排序: 时间复杂度O(n^2)，空间复杂度O(1),稳定排序
	BubbleSort(srcArray)

	// 插入排序: 时间复杂度O(n^2)，空间复杂度O(1),稳定排序
	InsertSort(srcArray)
}

func SelectionSort(a []int) {
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
	fmt.Println("SelectionSort:", a)
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
		// for j := 0; j < i; j++ {
		// 	if a[i] < a[j] {
		// 		a[i], a[j] = a[j], a[i]
		// 	}
		// }

		tmp:=a[i]
		k:=i-1
		for(k>=0&&a[k]>tmp){
			k--
		}
		for j:=i;j>k+1;j--{
			a[j] = a[j-1]  //向右移动
		}
	}
	fmt.Println("InsertSort:", a)
}

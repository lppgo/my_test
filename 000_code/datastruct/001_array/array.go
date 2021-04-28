package array

import "fmt"

func ExampleArray() {
	// 数组的定义
	var arr1 [10]int              // 定义长度为10的整形数组
	arr2 := [5]int{1, 2, 3, 4, 5} //定义并初始化,自动推导类型
	arr3 := [5]int{1, 2}          //指定总长度，前几位被初始化，没有的使用零值
	arr4 := [5]int{2: 10, 4: 11}  //有选择的初始化，没被初始化的使用零值
	arr5 := [...]int{2, 3, 4}     //自动计算长度
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3)
	fmt.Println("arr4:", arr4)
	fmt.Println("arr5:", arr5)

	// 数组的长度
	fmt.Println("len(arr3):", len(arr3))

	// 数组元素获取操作会引发类型的变化，数组将会转化为 Go 中新的数据类型切片Slice
	arr := arr2[:] // 代表所有元素
	fmt.Println("arr 1 :", arr)
	arr = arr2[:5] // 代表前五个元素，即区间的左闭右开
	fmt.Println("arr 2 :", arr)
	arr = arr[5:] // 代表从第5个开始（不包含第5个）
	fmt.Println("arr 3 :", arr)

	// 数组遍历方式一：for 循环遍历
	arr6 := [3]int{1, 2, 3}
	for i := 0; i < len(arr6); i++ {
		fmt.Println(arr6[i])
	}

	// 数组遍历方式二：for-range 遍历
	arr7 := [3]int{1, 2, 3}
	for k, v := range arr7 {
		fmt.Println(k) //元素位置
		fmt.Println(v) //元素值
	}

}

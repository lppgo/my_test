[TOC]
# 二： 一些整理的题目

### 1:  Slice

```go
// 数组切片的知识点
// 1: 基本结构
// 2: Slice扩容
func mySliceArray() {
	nums := [3]int{}
	nums[0] = 1

	fmt.Printf("nums: %v , len: %d, cap: %d\n", nums, len(nums), cap(nums)) //
	dnums := nums[0:2]
	dnums[0] = 5
	fmt.Printf("nums: %v ,len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("dnums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums))
	//fmt.Println(drums[2])
}
```

### 2:  copy()函数

```go
func myCopy() {
	dst := []int{1, 2, 3}
	src := []int{4, 5, 6, 7, 8}
	n := copy(dst, src)
	fmt.Printf("dst: %v, n: %d", dst, n)
}
```

### 3:  interface{} ，鸭子类型，简单工厂模式

### 5:  工厂模式

### 6:  值接收者和指针接收者

### 7：数组是值类型,切片是引用类型

### 8：切片反转

```go
func reverse() {
	s := []int{0, 1, 2, 3, 4, 5}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}
```

### 

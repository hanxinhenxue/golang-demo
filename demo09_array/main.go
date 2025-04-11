package main

import "fmt"

func main() {
	// 1. 数组的长度是类型的一部分，必须在声明时指定长度，没有长度的是切片
	var arr1 [3]int
	fmt.Printf("arr1: %T\n", arr1) // arr1: [3]int

	// 2.数组的初始化
	var arr2 [3]int

	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	// arr2[3] = 4 // 数组越界错误

	fmt.Println(arr2) // [1 2 3]

	// 自动根据元素个数推导数组长度
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arr3: %T\n", arr3) // arr3: [5]int

	// 指定下标
	var arr4 = [...]int{0: 1, 2: 3, 4: 5}
	fmt.Println(arr4) // [1 0 3 0 5]

	// 循环使用for range遍历数组
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	for key, val := range arr4 {
		fmt.Println(key, val)
	}

	// 数组是值类型，传递的是值的拷贝
	arr5 := [3]int{1, 2, 3}
	arr6 := arr5
	arr5[0] = 100
	fmt.Println(arr5, arr6) // [100 2 3] [1 2 3]

	// 多维数组
	arr7 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// arr7 := [...][3]int{{1, 2, 3}, {4, 5, 6}} // 也可以使用...来自动推导长度
	fmt.Println(arr7) // [[1 2 3] [4 5 6]]
}

package main

import "fmt"

func main() {
	// 切片是引用类型，使用copy来复制
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := sliceA
	sliceB[0] = 100
	fmt.Println(sliceA) // [100 2 3 4 5]
	fmt.Println(sliceB) // [100 2 3 4 5]

	// 使用copy函数复制切片
	sliceC := []int{1, 2, 3, 4, 5}
	sliceD := make([]int, len(sliceC)) // 创建一个和sliceC一样长度的切片
	copy(sliceD, sliceC)               // 复制sliceC到sliceD
	sliceD[0] = 100
	fmt.Println(sliceC) // [1 2 3 4 5]
	fmt.Println(sliceD) // [100 2 3 4 5]

}

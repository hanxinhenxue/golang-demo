package main

import "fmt"

func main() {
	// 使用make函数创建切片
	// make([]T, size, cap)
	a := make([]int, 4, 8)   // 创建一个长度为4，容量为8的切片
	fmt.Printf("a: %v\n", a) // a: [0 0 0 0]
	// 赋值通过下标 a[i] = value,但是无法通过下标进行扩容

	// 使用append函数扩容
	// append函数会返回一个新的切片，原来的切片不会改变
	var sliceA []int
	fmt.Printf("%v--%v--%v\n", sliceA, len(sliceA), cap(sliceA)) // []--0--0
	sliceA = append(sliceA, 1)                                   // 追加元素，扩容
	fmt.Printf("%v--%v--%v\n", sliceA, len(sliceA), cap(sliceA)) // [1]--1--1

	// 使用append合并切片
	sliceB := []int{2, 3, 4}
	sliceC := []int{5, 6, 7}
	sliceD := append(sliceB, sliceC...)                          // 合并切片，...表示拆分切片
	fmt.Printf("%v--%v--%v\n", sliceD, len(sliceD), cap(sliceD)) // [2 3 4 5 6 7]--6--6

}

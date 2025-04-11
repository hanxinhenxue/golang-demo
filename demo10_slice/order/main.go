package main

import "fmt"

func main() {
	// 从小到大排序 选择排序
	var numSlice = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				tmp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = tmp
			}
		}
	}
	fmt.Println(numSlice)

	// 冒泡排序
	var numSlice2 = []int{9, 8, 7, 5, 6, 4}
	for i := 0; i < len(numSlice2); i++ {
		for j := 0; j < len(numSlice2)-1; j++ {
			if numSlice2[j] > numSlice2[j+1] {
				tmp := numSlice2[j]
				numSlice2[j] = numSlice2[j+1]
				numSlice2[j+1] = tmp
			}
		}
	}
	fmt.Println(numSlice2)
}

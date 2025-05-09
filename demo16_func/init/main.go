package main

import "fmt"

// Int类型升序排序
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

// Int类型降序排序
func sortIntDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func main() {
	var sliceA = []int{12, 34, 37, 35, 556, 32, 2}
	arr := sortIntAsc(sliceA)
	fmt.Println(arr)
	arr2 := sortIntDesc(sliceA)
	fmt.Println(arr2)
}

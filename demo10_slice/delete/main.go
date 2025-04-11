package main

import "fmt"

func main() {
	// 切片没有专门的删除方法，删除元素可以使用append函数
	a := []int{1, 2, 3, 4, 5}
	// 删除索引2的元素，最后一个元素要加...
	a = append(a[:2], a[3:]...)
	fmt.Printf("%v\n", a) // [1 2 4 5]
}

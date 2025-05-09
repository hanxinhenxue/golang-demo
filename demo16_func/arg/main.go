package main

import "fmt"

// 求两个数的和
func sumFn(x int, y int) int {
	return x + y
}

// 函数参数简写，如果参数类型相同，可以省略类型
// 省略类型的参数必须在最后
func sumFn2(x, y int) int {
	return x + y
}

// 函数的可变参数
// 可变参数必须在最后
func sumFn3(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

// 函数可以一次性返回多个值
// 返回的类型需要括起来
func sumFn4(x int, y int) (int, int) {
	return x + y, x - y
}

// 将返回值声明 (sum, sub int)
func sumFn5(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}
func main() {
	sum1 := sumFn(1, 2)
	fmt.Println(sum1) // 3

	sum2 := sumFn2(1, 2)
	fmt.Println(sum2) // 3

	sum3 := sumFn3(1, 2, 3, 4, 5)
	fmt.Println(sum3) // 15

	a, b := sumFn4(1, 2)
	fmt.Println(a, b) // 3 -1
}

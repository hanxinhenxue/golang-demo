package main

import "fmt"

// 函数作为另一个函数的参数
func add(x, y int) int { // 定义一个函数 add，接收两个 int 类型的参数，返回它们的和
	return x + y
}
func sub(x, y int) int { // 定义一个函数 sub，接收两个 int 类型的参数，返回它们的差
	return x - y
}

type calcType func(int, int) int

func calc(a, b int, f calcType) int { // 定义一个函数 calc，接收三个参数 a, b, f
	return f(a, b) // 返回 f(a, b) 的结果
}
func main() {
	sum := calc(1, 2, add)
	fmt.Println("1 + 2 =", sum) // 输出 1 + 2 = 3
	diff := calc(1, 2, sub)
	fmt.Println("1 - 2 =", diff) // 1 - 2 = -1
}

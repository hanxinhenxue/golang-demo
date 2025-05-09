package main

import "fmt"

// 函数类型的定义和使用
type calc func(int, int) int // 定义一个函数类型 calc，接收两个 int 类型的参数，返回一个 int 类型的值

func add(x, y int) int { // 定义一个函数 add，接收两个 int 类型的参数，返回它们的和
	return x + y
}
func main() {
	var c calc
	c = add
	fmt.Printf("c的类型：%T\n", c) // 输出函数类型 main.calc

	f := add
	fmt.Printf("f的类型：%T\n", f) // 输出函数类型 func(int, int) int
}

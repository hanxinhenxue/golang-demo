package main

import "fmt"

func main() {
	/* fmt.Println("A Hello, World!") // 换行
	fmt.Print("B Hello, World!")   // 不换行
	fmt.Print("C Hello, World!")   // 不换行 */

	// fmt.Println("A", "B", "C") // A B C 同时输出多个中间有空格
	// fmt.Print("A", "B", "C")   // ABC 没有空格

	// a := "aaa"
	// fmt.Printf(a) // aaa

	var a int = 10
	var b int = 20
	c := 30 // c := 30 变量类型自动推导

	fmt.Printf("a=%v, b=%v, c=%v", a, b, c) // %v输出变量的值

	fmt.Printf("a=%v,a的类型是%T\n", a, a) // %T输出变量的类型
	// a=10,a的类型是int
}

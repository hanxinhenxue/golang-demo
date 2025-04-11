package main

import "fmt"

func main() {
	// goto语句可以跳转到指定标签处，类似于C语言的goto语句
	n := 30
	if n > 24 {
		fmt.Println("大于24")
		goto label1 // 跳转到label1标签处
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
label1:
	fmt.Println("ccc")
	// 打印 结果：
	// 大于24
	// ccc
}

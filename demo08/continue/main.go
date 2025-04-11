package main

import "fmt"

func main() {
	// continue可以跳过当前循环，继续下一次循环，仅限于for循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
			// continue label1 可以跳过指定标签的循环
			// continue label1 // 这里的label1是一个标签，表示跳过label1标签的循环
		}
		fmt.Println(i)
	}

}

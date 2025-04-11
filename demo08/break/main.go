package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
	// 结果：0 1

	/* i=0 j=0
	i=0 j=1
	i=0 j=2
	i=1 j=0
	i=1 j=1
	i=1 j=2  */
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break
			}
			fmt.Printf("i=%v j=%v \n", i, j)
		}
	}

	// 使用label跳出多层循环
	/* 	i=0 j=0
	i=0 j=1
	i=0 j=2  */
label1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break label1 // 跳出label1
			}
			fmt.Printf("i=%v j=%v \n", i, j)
		}
	}
}

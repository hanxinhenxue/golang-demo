package main

import "fmt"

func main() {
	// 1.整型和整型之间的转换
	/* 	var a int8 = 20
	   	var b int16 = 40
	   	fmt.Println(a + int8(b)) // 60 */

	// 2.浮点型和浮点型之间的转换
	/* var a float32 = 20
	var b float64 = 40
	fmt.Println(a + float32(b)) // 60 */

	// 3.浮整型和浮点型之间的转换
	var a float32 = 20.1
	var b int = 40
	fmt.Println(a + float32(b)) // 60.1

	// 转换低位像高位转换

}

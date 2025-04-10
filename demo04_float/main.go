package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// float32 float64两种浮点型 默认声明float64
	var f float32 = 3.14
	fmt.Printf("值：%v-----%f,类型%T\n", f, f, f) // 值：3.14-----3.140000,类型float32
	fmt.Println(unsafe.Sizeof(f))             // float32占用4个字节 float64占用8个字节

	// %f输出保留小数 %.2f保留两位小数
	var c float64 = 3.1415923
	fmt.Printf("值：%v---%f---%.2f\n", c, c, c) // 3.1415923---3.141592---3.14

	// float精度丢失
	d := 1129.6
	fmt.Println(d * 100) // 112959.99999999999

	// 使用decimal包来解决
}

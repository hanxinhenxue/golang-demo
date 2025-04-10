package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1. 声明整型
	// var a int = 10

	// int8 int16 int32 int64 有符号整型 第一位都是符号位+-
	// uint8 uint16 uint32 uint64 无符号整型，第一位最小是0

	// var num int8 = 130 // 超出范围
	// fmt.Println("numn=%v,类型：%T", num, num) // numn=130,类型:int8

	// unsafe.Sizeof(a)获取长度
	var a int8 = 15
	fmt.Printf("numn=%v,类型：%T\n", a, a) // numn=15,类型：int8
	fmt.Println(unsafe.Sizeof(a))       // 1 1个字节 int32 是4个字节

	var n1 uint8 = 255
	fmt.Printf("numn=%v,类型：%T\n", n1, n1) // numn=255,类型：uint8

	// int不同类型不能直接相加
	var a1 int32 = 10
	var a2 int64 = 20
	fmt.Println(int64(a1) + a2) // 30
	// fmt.Println(a1 + a2) // 不能直接相加
	// 高位转换低位可能是溢出，低位转高位可以

	num := 9
	fmt.Printf("num=%v,类型：%T\n", num, num) // num=30,类型：int
	fmt.Println(unsafe.Sizeof(num))        // 64位计算机 8个字节
	fmt.Printf("原样输出%v\n", num)            // 原样输出9
	fmt.Printf("2进制%b\n", num)             // 2进制1001
	fmt.Printf("8进制%o\n", num)             // 8进制11
	fmt.Printf("10进制%d\n", num)            // 10进制9
	fmt.Printf("16进制%x\n", num)            // 16进制9
}

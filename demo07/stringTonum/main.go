package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1.使用string类型转为整型

	str := "aaa"
	fmt.Printf("值: %v 类型: %T\n", str, str) // 值: 123456 类型: string

	/*
		ParseInt
		参数1：string数据
		参数2：进制，10表示十进制，16表示十六进制
		参数3：位数， 32， 64 16
		返回值：num err：转换失败返回0
	*/
	num, _ := strconv.ParseInt(str, 10, 64) // 10表示十进制，64表示int64类型
	fmt.Printf("值: %v, 类型: %T\n", num, num) // 值: 123456 类型: int64

	// 2.使用string类型转为bool 不建议
	str1 := "true"
	b, _ := strconv.ParseBool(str1)
	fmt.Printf("值: %v, 类型: %T\n", b, b) // 值: true, 类型: bool

}

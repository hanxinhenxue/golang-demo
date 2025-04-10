package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1.其他类型转换string类型
	// sprintf使用 int为%d float为%f bool为%t byte为%c

	var i int = 20
	var f float64 = 20.1
	var t bool = true
	var b byte = 'a'

	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("值: %v 类型: %T\n", str1, str1) // 值: 20 类型: string

	str2 := fmt.Sprintf("%f", f)
	fmt.Printf("值: %v 类型: %T\n", str2, str2) // 值: 20.100000 类型: string

	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("值: %v 类型: %T\n", str3, str3) // 值: true 类型: string

	str4 := fmt.Sprintf("%c", b)
	fmt.Printf("值: %v 类型: %T\n", str4, str4) // 值: a 类型: string

	// 2.使用strconv包转换string类型
	/*
		FormatInt
		参数1：int64类型的整数
		参数2：进制，10表示十进制，16表示十六进制
		返回值：string类型的整数
	*/
	var num1 int = 20
	s1 := strconv.FormatInt(int64(num1), 10)
	fmt.Printf("值: %v 类型: %T\n", s1, s1) // 值: 20 类型: string

	/*
		FormatFloat
		参数1：要转换的值
		参数2：格式化类型
		f 表示十进制浮点数
		b 表示二进制浮点数
		e 表示科学计数法
		E 表示科学计数法
		g
		G 表示根据值的大小自动选择格式
		参数3：精度，表示小数点后保留的位数 -1表示默认值
		参数4：64， 32

	*/
	var num2 float32 = 20.23
	s2 := strconv.FormatFloat(float64(num2), 'f', 2, 64)
	fmt.Printf("值: %v 类型: %T\n", s2, s2) // 值: 20.23 类型: string

	var num3 bool = true
	s3 := strconv.FormatBool(num3)
	fmt.Printf("值: %v 类型: %T\n", s3, s3) // 值: true 类型: string

	var num4 byte = 'a'
	s4 := strconv.FormatUint(uint64(num4), 10)
	fmt.Printf("值: %v 类型: %T\n", s4, s4) // 值: 97 类型: string

}

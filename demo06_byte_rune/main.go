package main

import "fmt"

func main() {
	// 1.字符属于int类型 单引号代表字符不是字符串
	var a = 'a'
	fmt.Printf("值：%v, 类型：%T\n", a, a) // 值：97, 类型：int32

	var str = "this"
	fmt.Printf("值：%v, 类型：%T，原样输出字符型：%c\n", str[2], str[2], str[2]) // 值：105, 类型：uint8 原样输出字符型：i

	// 一个汉字占用3个字节(utf-8)，
	// unsafe.Sizeof(a)无法查看string类型占据的空间，用len
	var b = '国'
	fmt.Printf("值：%v, 类型：%T，原样输出字符型：%c\n", b, b, b) // 值：22269, 类型：int32，原样输出字符型：国

	// 3.循环输出字符串，使用for相当于byte类型
	/* s := "hello 张三"
	// byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)\n", s[i], s[i])
	}

	// rune
	for _, v := range s {
		fmt.Printf("%v(%c)\n", v, v) // 直接输出字符
	} */

	// 4.字符串转byte和rune
	s1 := "big"
	// s1[0] = "p" // 报错
	bytes := []byte(s1) // 每一位是字母使用byte，如果有汉字使用rune
	// bytes := []rune(s1) // 每一位是汉字使用rune
	bytes[0] = 'p'
	fmt.Println(string(bytes)) // pig
}

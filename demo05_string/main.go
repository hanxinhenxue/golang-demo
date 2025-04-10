package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello"
	fmt.Printf("s1的值：%v,类型：%T\n", s1, s1) // s1的值：hello,类型：string

	// 使用\转义
	str2 := "hello\nworld"
	fmt.Println(str2)

	// 使用模板字符串输出多行字符串
	str3 := `this is line1
this is line2
this is line3`
	fmt.Println(str3)

	// 1. len(str)求长度
	str4 := "你好"
	fmt.Println(len(str4)) // 6，中文字符占用3个字节

	// 2.字符串拼接 + pringtf
	sTr1 := "你好"
	sTr2 := "世界"
	//(1): sTr3 := sTr1 + sTr2
	// (2)
	sTr3 := fmt.Sprintf("%v%v", sTr1, sTr2)
	fmt.Println(sTr3) // 你好世界

	// 3.分割字符串，需要strings包
	str5 := "hello,world,你好"
	arr := strings.Split(str5, ",")
	fmt.Println(arr) // [hello world 你好]

	// 4.切片转换字符串
	str6 := strings.Join(arr, "*")
	fmt.Println(str6) // hello*world*你好

	// 5.strings.contains判断一个字符串是否包含另个一字符串

	str7 := "hello world"
	str8 := "hello"
	fmt.Println(strings.Contains(str7, str8)) // true

	// 6.strings.HasPrefix判断一个字符串是否以另一个字符串开头
	// strings.HasSuffix判断一个字符串是否以另一个字符串结尾

	// 7.字符串出现的位置 strings.Index Index从前往后，LastIndex从后往前
	// 找不到返回-1
	str9 := "this is str"
	str10 := "is"
	fmt.Println(strings.Index(str9, str10))     // 2
	fmt.Println(strings.LastIndex(str9, str10)) //5
}

package main

import "fmt"

func main() {
	// switch使用 break可以省略
	ext := ".txt"

	// switch ext:= ".txt"; ext { // 这里的ext是局部变量，和上面的ext不是同一个变量
	switch ext {
	case ".txt":
		fmt.Println("文本文件")
		break // 这里的break可以省略
	case ".jpg", ".png": // 多个case可以合并
		// 这里的jpg和png是同一个case
		fmt.Println("图片文件")
	case ".mp3":
		fmt.Println("音乐文件")
	default:
		fmt.Println("未知文件")
	}

	// switch使用表达式示例
	// fallthrough可以穿透下一个case
	// 结果： 工作 退休
	age := 30
	switch {
	case age < 24:
		fmt.Println("学习")
	case age >= 24 && age < 60:
		fmt.Println("工作")
		fallthrough // fallthrough可以穿透下一个case
	case age > 60:
		fmt.Println("退休")
	default:
		fmt.Println("错误")
	}

}

package main

import "fmt"

func getInfo() (string, int) {
	return "用户一", 18
}

func main() {
	/* 1. var声明变量
	var 变量名称 类型 */

	var username string
	fmt.Println(username) // 默认值是空字符串

	var name string
	name = "张三"
	fmt.Println(name) // 张三

	var a1, a2 string
	a1 = "aaa"
	a2 = "bbb"
	fmt.Println(a1, a2) // aaa bbb

	// 一次声明多个变量
	var (
		// user string = "张三"
		// age = 18

		user string
		age  int
		sex  string
	)
	user = "张三"
	age = 18
	sex = "男"
	fmt.Println(user, age, sex)

	// 短变量声明只能作为局部变量 key := value
	// 不能在函数外部使用
	//  a, b, c := 1, 2, "C"

	// 匿名变量 _ 用于丢弃不需要的值
	// var _, userAge = getInfo()

	// const 同时声明多个常量，如果省略值表示和上面一样
	// const (
	// 	n1 = 1
	// 	n2
	// 	n3 = 2
	// 	n4
	// )
	// fmt.Println(n1, n2, n3, n4) // 1 1 2 2

	// iota每次自增1
	// const i0 = iota
	// fmt.Println(i0) // 0

	// const (
	// 	i1 = iota
	// 	i2
	// 	i3
	// 	i4
	// )
	// fmt.Println(i1, i2, i3, i4) // 0 1 2 3

	// 使用_跳过一个值
	// const (
	// 	i1 = iota
	// 	_
	// 	i3
	// 	i4
	// )
	// fmt.Println(i1, i3, i4) // 0 2 3

	// iota声明插队
	// const (
	// 	i1 = iota
	// 	i2 = 100
	// 	i3 = iota
	// 	i4
	// )
	// fmt.Println(i1, i2, i3, i4) // 0 100 2 3

	// 多个iota定义在一行
	const (
		n1, n2 = iota + 1, iota + 2
		n3, n4
		n5, n6
	)
	fmt.Println(n1, n2, n3, n4, n5, n6) // 1 2 2 3 3 4
}

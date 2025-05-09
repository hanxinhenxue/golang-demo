package main

import "fmt"

func main() {
	// 值类型：改变变量副本值的时候不会改变变量本身的值（基本类型、数组）
	// 引用类型：改变变量副本值的时候，只会改变变量本身的值（切片、map）

	var userInfo1 = make(map[string]string)
	userInfo1["username"] = "张三"
	userInfo1["age"] = "18"
	userInfo2 := userInfo1
	fmt.Println(userInfo1, userInfo2)
	userInfo2["username"] = "李四"
	fmt.Println(userInfo1, userInfo2)
}

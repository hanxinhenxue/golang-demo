package main

import "fmt"

func main() {
	// 1.map初始化的方法
	var userInfo = make(map[string]string)

	userInfo["name"] = "zhangsan"
	userInfo["age"] = "18"
	userInfo["sex"] = "man"

	fmt.Println(userInfo) //map[age:18 name:zhangsan sex:man]

	// 2.map初始化的方法
	userInfo1 := map[string]string{
		"name": "lisi",
		"age":  "19",
		"sex":  "woman",
	}
	fmt.Println(userInfo1)

	// 循环map
	for k, v := range userInfo1 {
		fmt.Println(k, v)
	}

	// 查找map
	val, isOk := userInfo["name"]
	fmt.Println(val, isOk) // zhangsan true

	// 删除mapd元素
	delete(userInfo, "name")
	fmt.Println(userInfo)
}

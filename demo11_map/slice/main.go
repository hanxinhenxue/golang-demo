package main

import "fmt"

func main() {
	// 1.定义一个map类型的切片
	var userInfo = make([]map[string]string, 2, 2)

	// map不初始化默认值是nil
	fmt.Println(userInfo[0] == nil) // true

	if userInfo[0] == nil {
		userInfo[0] = make(map[string]string)
		userInfo[0]["username"] = "张三"
		userInfo[0]["age"] = "20"
		userInfo[0]["height"] = "180cm"
	}

	if userInfo[1] == nil {
		userInfo[1] = make(map[string]string)
		userInfo[1]["username"] = "李四"
		userInfo[1]["age"] = "22"
		userInfo[1]["height"] = "170cm"
	}

	fmt.Println(userInfo)

	// 循环遍历
	for _, v := range userInfo {
		for k, v := range v {
			fmt.Println(k, v)
		}
	}

	// 2.定义一个切片类型的map
	var userInfo1 = make(map[string][]string)
	userInfo1["hobby"] = []string{"吃饭", "睡觉"}
	fmt.Println(userInfo1)
}

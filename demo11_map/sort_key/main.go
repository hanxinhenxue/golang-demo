package main

import (
	"fmt"
	"sort"
)

// map类型键升序排列
func mapSort(map1 map[string]string) string {
	var sliceKey []string
	for k, _ := range map1 {
		sliceKey = append(sliceKey, k)
	}
	sort.Strings(sliceKey)
	var str string
	for _, k := range sliceKey {
		str += fmt.Sprintf("%v=>%v ", k, map1[k])
	}
	return str
}

func main() {
	m1 := map[string]string{
		"username": "张三",
		"password": "123456",
		"age":      "18",
		"address":  "北京",
	}
	fmt.Println(mapSort(m1)) // address=>北京 age=>18 password=>123456 username=>张三
}

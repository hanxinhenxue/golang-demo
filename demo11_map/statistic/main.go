package main

import (
	"fmt"
	"strings"
)

func main() {
	// 统计一个字符串字符出现的次数
	str := "how do you do"
	var strSlice = strings.Split(str, " ")

	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}

package main

import "fmt"

func main() {
	// for range使用
	str := "你好golang"
	for key, val := range str {
		fmt.Printf("key: %v, val: %v\n", key, val)
	}

	arr := []string{"你好", "golang"}
	for key, val := range arr {
		fmt.Printf("key: %v, val: %v\n", key, val)
	}
}

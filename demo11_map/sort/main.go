package main

import (
	"fmt"
	"sort"
)

func main() {
	// map打印的时候，它的顺序是随机的
	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 13
	map1[8] = 90
	map1[4] = 56

	for k, v := range map1 {
		fmt.Println(k, v)
	}
	/* 	10 100
	1 13
	4 56
	8 90 */

	// map排序
	// 1. 将map的key放入切片中
	// 2. 对切片进行排序
	var keys []int
	for k := range map1 {
		keys = append(keys, k)
	}
	// 3. 对切片进行排序
	sort.Ints(keys)
	// 4.循环遍历map
	for _, k := range keys {
		fmt.Println(k, map1[k])
	}
	/* 	1 13
	4 56
	8 90
	10 100 */
}

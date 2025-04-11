package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort包升序排序
	intList := []int{2, 3, 4, 5, 7, 8, 9, 1, 6}
	float8List := []float64{2.3, 3.4, 4.5, 5.6, 7.8, 8.9, 1.2, 6.7}
	stringList := []string{"b", "a", "g", "e", "c", "k", "d", "h", "f"}

	sort.Ints(intList)        // 升序排序
	sort.Float64s(float8List) // 升序排序
	sort.Strings(stringList)  // 升序排序
	fmt.Println(intList)      // [1 2 3 4 5 6 7 8 9]
	fmt.Println(float8List)   //[1.2 2.3 3.4 4.5 5.6 6.7 7.8 8.9]
	fmt.Println(stringList)   //[a b c d e f g h k]

	// sort包降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println(intList)    //[9 8 7 6 5 4 3 2 1]
	fmt.Println(float8List) //[8.9 7.8 6.7 5.6 4.5 3.4 2.3 1.2]
	fmt.Println(stringList) //[k h g f e d c b a]
}

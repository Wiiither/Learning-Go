package main

import "fmt"

// 范围
func main() {
	//  范围用于遍历数组、切片、字符串等

	//  定义一个范围
	var range1 = []int{1, 2, 3}
	for i := range range1 {
		fmt.Println("range1[", i, "] = ", range1[i])
	}

	//  范围的索引和值
	for i, v := range range1 {
		fmt.Println("range1[", i, "] = ", v)
	}
}

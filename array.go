package main

import "fmt"

func main() {
	//  数组
	//  var 数组名 [长度]类型
	var arr [10]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println("arr = ", arr) //  arr =  [1 2 0 0 0 0 0 0 0 0]

	//  数组的长度
	//  数组的长度
	fmt.Println("len(arr) = ", len(arr)) //  len(arr) =  10

	//  如果长度不指定，则数组长度为 0
	var arr2 []int
	fmt.Println("arr2 = ", arr2) //  arr2 =  []

	//  如果长度不确定，则可以使用 ... 来表示
	var arr3 = [...]int{1, 2, 3}
	fmt.Println("arr3 = ", arr3) //  arr3 =  [1 2 3]

	//  数组的遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println("arr3[", i, "] = ", arr3[i])
	}

	//  数组的切片
	var arr4 = arr3[1:3]
	fmt.Println("arr4 = ", arr4) //  arr4 =  [2 3]

	//  数组的切片
	var arr5 = arr3[:3]
	fmt.Println("arr5 = ", arr5) //  arr5 =  [1 2 3]

	//  数组的切片
	var arr6 = arr3[1:]
	fmt.Println("arr6 = ", arr6) //  arr6 =  [2 3]

	//  数组的切片
	var arr7 = arr3[:]
	fmt.Println("arr7 = ", arr7) //  arr7 =  [1 2 3]
}

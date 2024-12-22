package main

import "fmt"

func main() {
	//  切片
	//  Go 数组的长度不可改变
	//  切片是数组的抽象
	//  切片是引用类型
	//  切片的长度可以改变
	//  切片是一个指向数组的指针

	//  定义一个切片
	var slice []int
	slice = []int{1, 2, 3}
	fmt.Println("slice = ", slice) //  slice =  [1 2 3]

	//  切片的遍历
	for i := 0; i < len(slice); i++ {
		fmt.Println("slice[", i, "] = ", slice[i])
	}

	//  切片的切片
	var slice2 = slice[1:3]
	fmt.Println("slice2 = ", slice2) //  slice2 =  [2 3]

	//  切片的切片
	var slice3 = slice[:3]
	fmt.Println("slice3 = ", slice3) //  slice3 =  [1 2 3]

	//  切片的切片
	var slice4 = slice[1:]
	fmt.Println("slice4 = ", slice4) //  slice4 =  [2 3]

	//  切片的切片
	var slice5 = slice[:]
	fmt.Println("slice5 = ", slice5) //  slice5 =  [1 2 3]

	//  len 和 cap
	fmt.Println("len(slice) = ", len(slice)) //  len(slice) =  3
	fmt.Println("cap(slice) = ", cap(slice)) //  cap(slice) =  3
	//  它们的区别在于
	//  len 是切片的长度
	//  cap 是切片的容量
	//  容量是指切片底层数组的长度
	//  长度是指切片中元素的个数

	//  切片的追加
	slice = append(slice, 4)
	fmt.Println("slice = ", slice) //  slice =  [1 2 3 4]

	//  切片的拷贝
	var slice6 = make([]int, 3)
	copy(slice6, slice)
	fmt.Println("slice6 = ", slice6) //  slice6 =  [1 2 3]

	//  空切片
	//  一个切片在未初始化之前默认为 nil，长度为 0
	var slice7 []int
	fmt.Println("slice7 = ", slice7) //  slice7 =  []
}

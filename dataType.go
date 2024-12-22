package main

import "fmt"
import "reflect"

// 数据类型
func main() {
	aBool := false
	checkTypeWithReflect(aBool)

	//  整型
	aInt := 100
	checkTypeWithReflect(aInt)

	//  浮点型
	aFloat := 100.123
	checkTypeWithReflect(aFloat)

	//  字符串
	aString := "Hello, World!"
	checkTypeWithReflect(aString)

	//  数组
	aArray := [3]int{1, 2, 3}
	checkTypeWithReflect(aArray)

	//  切片
	aSlice := []int{1, 2, 3}
	checkTypeWithReflect(aSlice)

	//  字典
	aMap := map[string]int{"a": 1, "b": 2, "c": 3}
	checkTypeWithReflect(aMap)

	//  结构体
	aStruct := struct {
		Name string
		Age  int
	}{"张三", 18}
	checkTypeWithReflect(aStruct)

	//  指针
	aPointer := &aInt
	checkTypeWithReflect(aPointer)

	//  函数
	aFunction := func() {
		fmt.Println("这是一个函数")
	}
	checkTypeWithReflect(aFunction)

	//  通道
	aChannel := make(chan int)
	checkTypeWithReflect(aChannel)

	//  类型转换
	var a int = 10
	var b float64 = float64(a)
	fmt.Println("b = ", b) //  b =  10

	//  类型断言
	//  类型断言用于将一个接口类型转换为具体的类型
	//  类型断言的语法是：类型断言表达式.(类型)
	//  类型断言表达式可以是接口类型变量，也可以是接口类型变量.(类型)
	var a interface{} = 10
	var b int = a.(int)
	fmt.Println("b = ", b) //  b =  10

}

func checkTypeWithReflect(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("类型是:", t)
}

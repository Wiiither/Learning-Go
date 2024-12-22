package main

import "fmt"

func main() {
	//  接口
	//  接口是 Go 语言中的抽象类型
	//  接口是引用类型
	//  接口是实现多态的机制
	//  接口是 Go 语言中的强大工具
	//  接口可以用于实现多态、抽象、封装等

	//  定义一个接口
	type Animal interface {
		eat()
	}

	//  实现接口
	type Dog struct {
		name string
	}
	func (d Dog) eat() {
		fmt.Println("Dog eat")
	}

	//  接口的类型断言
	var a Animal = Dog{"dog"}
	var b Dog = a.(Dog)
	fmt.Println("b = ", b)	//  b =  {dog}

	//  类型选择
	var a interface{} = 10
	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	}	

	//  接口组合
	type Animal interface {
		eat()
	}
	type Bird interface {
		fly()
	}
	type Duck interface {
		Animal
		Bird
	}

	//  动态值和动态类型
	var a interface{} = 10
	fmt.Println("a = ", a)	//  a =  10
	fmt.Println("a.(int) = ", a.(int))	//  a.(int) =  10
}

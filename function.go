package main

import "fmt"

func main() {
	//  函数
	func add(a int, b int) int {
		return a + b
	}

	//  函数可以返回多个值
	func swap(a int, b int) (int, int) {
		return b, a
	}


	//  函数可以作为参数传递
	type FuncType func(int, int) int
	func test(a int, b int, f FuncType) int {
		return f(a, b)
	}

	//  函数可以作为返回值
	func test() FuncType {
		return add
	}	

	//  匿名函数
	func max(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	getMax := max()
	fmt.Println("getMax = ", getMax(1, 2))	//  getMax =  2

	//  方法
	//  方法是包含了接受者的函数
	type Person struct {
		name string
		age int
	}
	func (p Person) getName() string {
		return p.name
	}

	//  方法可以作为参数传递
	func test(p Person, f func(Person) string) string {
		return f(p)
	}

	//  方法可以作为返回值
	func test() func(Person) string {
		return getName
	}


	//  递归函数
	func fibonacci(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n - 1) + fibonacci(n - 2)
	}
}
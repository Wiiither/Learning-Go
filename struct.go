package main

import "fmt"

// 结构体
func main() {

	//  定义一个结构体
	type Person struct {
		name string
		age  int
	}

	//  定义一个结构体变量
	var person Person
	person.name = "Wither"
	person.age = 20
	fmt.Println("person = ", person) //  person =  {Wither 20}

	//  结构体指针
	var personPtr *Person = &person
	fmt.Println("personPtr = ", personPtr)                 //  personPtr =  0xc0000180a0
	fmt.Println("(*personPtr).name = ", (*personPtr).name) //  (*personPtr).name =  Wither
	fmt.Println("personPtr.name = ", personPtr.name)       //  personPtr.name =  Wither
}

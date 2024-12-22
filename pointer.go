package main

import "fmt"

func main() {
	//  指针
	var a int = 10
	var ptr *int = &a
	fmt.Println("ptr = ", ptr)   //  ptr =  0xc0000180a0
	fmt.Println("*ptr = ", *ptr) //  *ptr =  10

	//  空指针
	var ptrNull *int
	fmt.Println("ptrNull = ", ptrNull) //  ptrNull =  <nil>
}

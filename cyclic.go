package main

import "fmt"

// 循环语句
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	// for 循环的另一种形式
	var j int = 0
	for j < 10 {
		fmt.Println("j = ", j)
		j++
	}

	// for 循环的第三种形式
	var m int = 0
	for {
		fmt.Println("m = ", m)
		m++
		if m >= 10 {
			break
		}
	}

}

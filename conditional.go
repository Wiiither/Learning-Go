package main

import "fmt"

//  条件语句
func main() {
		//  条件语句
	var a int = 10
	if a > 0 {
		fmt.Println("a > 0")
	} else {
		fmt.Println("a <= 0")
	}
	
	switch a {
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	default:
		fmt.Println("a = 0")
	}

	select {		
	//  类似于 switch 语句，但是 select 语句只能用于通道的选择
	//  每个 case 都必须是一个通道操作
	//  所有通道表达式都会被求值
	//  所有被发送的表达式都会被求值
	//  如果任意一个通信可以进行，它就执行，其他被忽略
	//  如果有多个 case 都可以运行，Select 会随机公平地选出一个执行
	case a > 0:
		fmt.Println("a > 0")
	case a < 0:
		fmt.Println("a < 0")
	default:
		fmt.Println("a = 0")
	}
	//  ==================
	//  Go 语言没有三目运算符
	//  ==================
}
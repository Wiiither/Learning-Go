package main

import "fmt"

func main() {
	//  通道
	//  通道是连接多个协程的管道
	//  通道是引用类型
	//  通道是线程安全的
	//  通道是阻塞的

	//  定义一个通道
	var ch = make(chan int)

	//  通道的接收
	var ch = make(chan int)
	ch <- 1
	fmt.Println("ch = ", <-ch) //  ch =  1

	//  通道的遍历
	for i := range ch {
		fmt.Println("ch[", i, "] = ", ch[i])
	}

	//  通道的关闭
	close(ch)
}

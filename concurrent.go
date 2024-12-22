package main

import "fmt"

func main() {
	//  并发
	//  并发是 Go 语言中的并发机制
	//  通过 goroutine 和 channel 实现并发
	//  goroutine 是 Go 语言中的轻量级线程
	//  channel 是 Go 语言中的管道
	//  Scheduler 是 Go 语言中的调度器

	//  举一些例子
	go hello() {
		fmt.Println("Hello, World!")
	}()
		
	//  channel 的例子
	ch := make(chan int)
	go func() {
		ch <- 1		//  把 1 发送到 channel 中
	}()
	fmt.Println("ch = ", <- ch)	//  ch =  1

	//  通道缓冲区
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("ch = ", <- ch)	//  ch =  1

	//  通道的关闭
	close(ch)

	//  通道的遍历
	for i := range ch {
		fmt.Println("ch[", i, "] = ", ch[i])
	}
}

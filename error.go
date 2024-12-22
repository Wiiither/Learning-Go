package main

import "fmt"

func main() {
	//  错误处理
	//  错误处理是 Go 语言中的错误处理机制
	//  主要围绕以下机制展开
	//  1. error 接口：标准的错误表示
	//  2. 显示返回值：通过函数的返回值返回错误
	//  3. 自定义错误：可以通过标准库或自定义的方式创建错误
	//  4. panic 和 recover：用于处理运行时错误

	//  1. error 接口
	//  error 接口是 Go 语言中的标准错误表示
	//  error 接口是接口类型
	//  error 接口是引用类型
	//  error 接口是空接口
	//  error 接口是标准库中的错误表示

	type error interface {
		Error() string
	}

	//  用 errors 包创建错误
	var err error = errors.New("这是一个错误")
	fmt.Println("err = ", err)	//  err =  这是一个错误

	//  2. 显示返回值
	func test() (int, error) {
		return 1, errors.New("这是一个错误")
	}

	//  3. 自定义错误
	type MyError struct {
		msg string
	}
	func (e MyError) Error() string {
		return e.msg
	}

	//  errors.ls 和 errors.As
	//  从 Go 1.13 开始，errors 包提供了两个函数：errors.Is 和 errors.As
	//  errors.Is 用于判断错误是否相等
	//  errors.As 用于判断错误是否是指定的错误类型

	//  举个例子
	var err error = MyError{"这是一个错误"}
	fmt.Println("errors.Is(err, MyError) = ", errors.Is(err, MyError))	//  errors.Is(err, MyError) =  true
	fmt.Println("errors.As(err, &MyError) = ", errors.As(err, &MyError))	//  errors.As(err, &MyError) =  true

	//  4. panic 和 recover
	//  panic 用于抛出运行时错误
	//  recover 用于捕获运行时错误
	//  panic 和 recover 是 Go 语言中的强大工具
	//  举个例子
	func test() {
		defer func() {
			fmt.Println("recover = ", recover())	//  recover =  <nil>
		}()
		panic("这是一个错误")
	}
	test()
}

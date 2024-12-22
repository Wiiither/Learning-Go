// 定义了包名，必须在源文件中非注释的第一行指明这个文件属于哪个包
// package main 表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包
package main

//  fmt 包实现了格式化 IO（输入/输出）的函数
import "fmt"

// main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一次执行的函数（入口函数）
func main() { //  { 不能单独放在一行
	fmt.Println("Hello, World!" + "你好，世界！")

	//  Sprintf 根据格式化参数生成格式化字符串并返回该字符串
	var str = fmt.Sprintf("Hello, World! %d", 100)
	fmt.Println(str)
}

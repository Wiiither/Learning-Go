package main

import "fmt"

//  变量和常量
func main() {
	//  go 变量声明的方式 
	//  var 变量名 变量类型 = 变量值
	//  如果没有初始化，则变量默认为 0 值
	//  * 如果变量声明了没有使用，则编译不通过
	var age int = 100
	fmt.Println("age = ", age)
	
	var ageNull int
	fmt.Println("ageNull = ", ageNull)	//  ageNull =  0

	//  也可以省略类型
	var name = "Wither"
	fmt.Println("name = ", name)	//  name =  Wither	

	//  也可以省略 var 关键字
	sex := "男"
	fmt.Println("sex = ", sex)	//  sex =  男

	//  可以通过 & 符号获取变量的地址
	fmt.Println("&age = ", &age)	//  &age =  0xc0000180a0

	//  可以通过 * 符号获取变量的值
	fmt.Println("*age = ", *age)	//  *age =  100

	//  值类型变量通常存储在栈中，尤其是当他们是局部变量时。
	//  当值类型变量需要在函数作用于之外使用时, Go 会将其分配到堆内存中。

	//  常量，使用 const 进行声明
	const PI = 3.14
	fmt.Println("PI = ", PI)	//  PI =  3.14

	//  常量可以作为枚举使用
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	//  iota 是特殊常量，可以理解为 const 语句块中的行索引
	const (
		a = iota	//  0
		b = iota	//  1
		c = iota	//  2
	)

	const (
		a = iota // a = 0
		b        // b = 1 (iota 自动递增)
		c        // c = 2 (iota 自动递增)
		d = "ha" // d = "ha" (iota 不参与字符串常量的赋值)
		e        // e = "ha" (与 d 相同，使用上一个值)
		f = 100  // f = 100
		g        // g = 100 (与 f 相同，使用上一个值)
		h = iota // h = 1 (iota 重新开始计数)
		i        // i = 2 (iota 自动递增)
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)	//  0 1 2 ha ha 100 100 1 2
}

package main

import "fmt"

//  运算符
func main() {
	//  算术运算符
	var a int = 10
	var b int = 20
	fmt.Println("a + b = ", a + b)	//  a + b =  30
	fmt.Println("a - b = ", a - b)	//  a - b =  -10
	fmt.Println("a * b = ", a * b)	//  a * b =  200
	fmt.Println("a / b = ", a / b)	//  a / b =  0
	fmt.Println("a % b = ", a % b)	//  a % b =  10
	fmt.Println("a++ = ", a++)	//  a++ =  10
	fmt.Println("a-- = ", a--)	//  a-- =  11
	fmt.Println("++a = ", ++a)	//  ++a =  12
	fmt.Println("--a = ", --a)	//  --a =  11

	//  关系运算符
	fmt.Println("a == b = ", a == b)	//  a == b =  false
	fmt.Println("a != b = ", a != b)	//  a != b =  true
	fmt.Println("a > b = ", a > b)	//  a > b =  false
	fmt.Println("a < b = ", a < b)	//  a < b =  true
	fmt.Println("a >= b = ", a >= b)	//  a >= b =  false
	fmt.Println("a <= b = ", a <= b)	//  a <= b =  true

	//  逻辑运算符
	fmt.Println("a && b = ", a && b)	//  a && b =  false
	fmt.Println("a || b = ", a || b)	//  a || b =  true
	fmt.Println("!a = ", !a)	//  !a =  false

	//  位运算符
	fmt.Println("a & b = ", a & b)	//  a & b =  0
	fmt.Println("a | b = ", a | b)	//  a | b =  30
	fmt.Println("a ^ b = ", a ^ b)	//  a ^ b =  30
	fmt.Println("a << b = ", a << b)	//  a << b =  320
	fmt.Println("a >> b = ", a >> b)	//  a >> b =  0

	//  位运算符的主要作用是进行位运算，比如进行权限管理
	//  位运算符 & 按位与运算符：参与运算的两个值，如果两个相应位都为1，则该位的结果为1，否则为0
	//  位运算符 | 按位或运算符：参与运算的两个值，如果两个相应位有一个为1，则该位的结果为1，否则为0
	//  位运算符 ^ 按位异或运算符：参与运算的两个值，如果两个相应位相同，则该位的结果为0，否则为1
	//  位运算符 << 左移运算符：将参与运算的值的二进制位向左移动，高位丢弃，低位补0
	//  位运算符 >> 右移运算符：将参与运算的值的二进制位向右移动，低位丢弃，高位补0

	//  赋值运算符
	var a int = 10
	var b int = 20
	a = b
	fmt.Println("a = ", a)	//  a =  20

	//  复合赋值运算符
	a += b
	fmt.Println("a = ", a)	//  a =  30

	a -= b
	fmt.Println("a = ", a)	//  a =  10
	a *= b
	fmt.Println("a = ", a)	//  a =  200
	a /= b
	fmt.Println("a = ", a)	//  a =  10
	a %= b
	fmt.Println("a = ", a)	//  a =  10

	//  其他运算符
	var a int = 10
	var ptr *int = &a
	fmt.Println("ptr = ", ptr)	//  ptr =  0xc0000180a0
	fmt.Println("*ptr = ", *ptr)	//  *ptr =  10	
}

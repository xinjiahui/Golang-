package main

import "fmt"

func main() {
	var (
		a = 2
		b = 5
	)
	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在=右边赋值 a++ ==> a = a+1
	b--
	fmt.Println(a)
	//关系运算符,go语言是强类型，相同类型的变量才能比较
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
	//逻辑运算符
	// 如果年龄大于18 并且 年龄小于60所
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("上班")
	} else {
		fmt.Println("不用上班")
	}
	//如果年龄小于18岁 或者 年龄大于60岁
	if age < 18 || age > 60 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("上班")
	}
	//not取反，原来为真就为假
	isMarried := false
	fmt.Println(isMarried)  //false
	fmt.Println(!isMarried) //true
	// 位运算：针对的是二进制数
	// &： 安位与，2位都是1，为1
	fmt.Println(5 & 2) //000
	// |:按位或,两位有一个为1就为1
	fmt.Println(5 | 2) // 111
	// ^： 按位异或（两位不一样的则为1）
	fmt.Println(5 ^ 2) // 111
	// <<: 将二进制位左移指定位数
	fmt.Println(5 << 1)  // 1010-> 10
	fmt.Println(1 << 10) // 10000000000-> 1024
	// >> :将二进制位右移指定位数
	fmt.Println(5 >> 1) // 10->2
	var m = int8(1)
	fmt.Println(m << 10) //1024,不建议
	// 赋值运算，用来给变量赋值的
	var x int
	x = 10
	x += 1  //x = x + 1
	x -= 1  //x = x - 1
	x *= 2  //x = x * 2
	x /= 2  //x = x / 2
	x %= 2  //x = x % 2
	x <<= 2 //  x = x << 2
	x &= 2  // x = x & 2
	x |= 3  // x = x | 3
	x ^= 4  //x = x ^ 4
	x >>= 2 // x = x >> 21
}

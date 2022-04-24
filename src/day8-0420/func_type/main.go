package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("hello shanghai")
}
func f2() int {
	return 10
}

func f4(x, y int) int {
	return x + y
}

//函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

//函数也可以作为返回值
func ff(a, b int) int {
	return a + b
}
func f5(x func() int) func(int, int) int {
	return ff
}
func main() {
	// a 和 b 是不同类型b
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(f2)
	f3(b)
	//f3(f4)
	f7 := f5(f2)
	fmt.Printf("%T\n", f7) //打印的是内存地址
}

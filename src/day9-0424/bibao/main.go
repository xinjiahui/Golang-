package main

import "fmt"

// 闭包是什么？
// 闭包是一个函数，这个函数包含了他外部作用域的一个变量

//底层原理
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找
func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 要求：
// f1(f2)

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		//fmt.Println(x +y)
		// f2()
		f(x, y)

	}
	return (tmp)
}

func main() {
	//ret := adder()
	//ret2 := ret(200)
	//fmt.Println(ret2)
	ret := f3(f2, 100, 200) // 把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	fmt.Printf("%T\n", ret)
	f1(ret)
}

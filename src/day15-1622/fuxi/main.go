package main

import "fmt"

//高阶函数：函数可以作为参数也可以作为返回值
func yuanshuai(name string) {
	fmt.Println("Hello", name)
}

//函数作为参数
func lixiang(f func(string), name string) {
	f(name)
}

// 函数作为返回值
func zhoulin() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func low(f func()) {
	f()
}

//闭包：函数和其外部变量的引用
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

//defer :延迟调用，多用于处理资源释放

//内置函数： panic和recover
func main() {
	lixiang(yuanshuai, "xinjiahui")
	ret := zhoulin()
	sum := ret(10, 20)
	fmt.Printf("%T\n", ret)
	fmt.Println(sum)
	fc := bi(yuanshuai, "元帅")
	low(fc)
}

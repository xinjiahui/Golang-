package main

import "fmt"

//给自定义类型加方法
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {

	m := myInt(100)
	fmt.Println(m)
	m.hello()

}

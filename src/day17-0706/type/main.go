package main

import "fmt"

//自定义类型和类型别名
//type后面跟的是类型

type myInt int     //自定义别名，编译完之后也保留
type yourInt = int //类型别名
func main() {
	var n myInt
	var m yourInt
	n = 100
	m = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	//rune 32位有符号整形，int32的别名，二者视为同一类型。
	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}

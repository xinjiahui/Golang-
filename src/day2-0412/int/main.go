package main

import "fmt"

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1) //10
	fmt.Printf("%o\n", i1) // 8
	fmt.Printf("%x\n", i1) //16
	fmt.Printf("%b\n", i1) //2

	i2 := 077
	fmt.Printf("%d\n", i2)
	//fmt.Sprintf()
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	// 查看变量类型
	fmt.Printf("%T\n", i1)

	// 声明int8 类型的变量
	i4 := int8(9) //明确指定int8类型，否则默认为int类型
	fmt.Printf("%T\n", i4)
}

package main

import "fmt"

// pointer

func main() {
	//1.&：取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) //*int:int类型指针
	//2.*:根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	//var a1 *int //// nil pointer
	var a = new(int)
	// new 函数申请一个内存地址
	*a = 100
	fmt.Println(*a)

	// make和new都是用来申请内存的
	//new很少用，一般用来给基本数据类型申请内存，string\int,返回的是对应类型的指针{*string *int}
	//make 是用来给slice，map,chan申请内存的，make函数返回的是对应的这三个类型本身

}

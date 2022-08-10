package main

import "fmt"

//结构体嵌套
type address struct {
	provice string
	city    string
}

type person struct {
	name string
	age  int
	addr address
	//address //匿名嵌套结构体
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "xinjiahui",
		age:  18,
		addr: address{
			provice: "黑龙江",
			city:    "大庆",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.addr.city)
	//fmt.Println(p1.city) //匿名的时候可以直接访问
}

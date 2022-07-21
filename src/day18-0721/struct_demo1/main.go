package main

import "fmt"

//结构体

type persion struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个persion类型的变量
	var p persion
	//通过字段赋值
	p.name = "xinjiahui"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球"}
	//访问变量p的字段
	fmt.Printf("%T\n", p)
	fmt.Println(p)
	var p2 persion
	p2.name = "test"
	p2.age = 18
	fmt.Printf("type: %T value: %v\n", p2, p2)
}

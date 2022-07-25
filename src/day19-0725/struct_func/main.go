package main

import "fmt"

//构造函数，返回一个结构体变量的函数
//结构体是值类型，赋值的时候都是拷贝
type persion struct {
	name string
	age  int
}

//构造函数
//返回的是结构体还是结构体指针
//当结构体比较大的时候，尽量使用结构体指针，减少程序的内存开销
func newPersion(name string, age int) *persion {
	return &persion{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPersion("xinjiahui", 18)
	p2 := newPersion("haha", 1999)
	fmt.Println(p1, p2)
}

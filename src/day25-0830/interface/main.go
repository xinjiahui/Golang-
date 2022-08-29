package main

import "fmt"

//接口是一种类型
//基本数据类型
//结构体类型（可以放各个维度的数据）
// 引出接口的实例

type cat struct{}
type dog struct{}
type persion struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵～")
}
func (d dog) speak() {
	fmt.Println("汪汪汪～")
}
func (p persion) speak() {
	fmt.Println("啊～")
}

func da() {
	//接收一个参数，传进来什么我就打什么
	x.speak() //挨打了就要叫
}
func main() {
	var c1 cat
	var d1 dog
	var p1 persion

	da(c1)
	da(d1)
	da(p1)
}

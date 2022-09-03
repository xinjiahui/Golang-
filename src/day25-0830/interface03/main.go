package main

import "fmt"

//接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("激动")
}
func (c chicken) eat(food string) {
	fmt.Printf("吃鸡%s...\n", food)
}
func (c cat) move() {
	fmt.Println("走猫步")
}
func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 animal
	fmt.Printf("%T\n", a1)
	bc := cat{
		name: "淘气",
		feet: 2,
	}
	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)
	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	fmt.Printf("%T\n", a1) //接口保存的分为两部分，值的类型和值本身，interface是一种抽象的类型
}

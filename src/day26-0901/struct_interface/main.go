package main

import "fmt"

//同一个结构体可以实现多个接口
//接口也可以嵌套
type mover interface {
	move()
}
type eater interface {
	eat(string)
}
type cat struct {
	name string
	feet int8
}
type animal interface {
	mover
	eater
}

//cat实现了mover接口

func (c *cat) move() {
	fmt.Println("走猫步...")
}

//cat 实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 eater
	c1 := &cat{"tom", 4}
	a1 = c1
	fmt.Println(a1)

}

package main

import "fmt"

//接口示例2

//不管是什么牌子的车，都能跑
//定义一个car接口类型
//不管什么结构体类型，只要有run方法，都是car 类型
type car interface {
	run()
}

type falali struct {
	brand string
}
type baoshijie struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s速度70迈~\n", f.brand)
}
func (b baoshijie) run() {
	fmt.Printf("%s速度是700迈~\n", b.brand)
}

//drive函数是接收一个car类型的变量
func drive(c car) {
	c.run()
}
func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)

}

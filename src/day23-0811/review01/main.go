package main

import "fmt"

//复习结构体

//有名字的结构体
type tmp struct {
	x int
	y int
}
type person struct {
	name string
	age  int
}

func sum(x int, y int) int {
	return x + y

}
func sum1(x int, y int) (ret int) {
	ret = x + y
	return ret
}

//构造（构造结构体变量的）函数,返回值就是对应的结构体类型
func newPerson(n string, i int) person {
	return person{
		name: n,
		age:  i,
	}
}

func newPerson1(n string, i int) (p person) {
	p = person{
		name: n,
		age:  i,
	}
	return p
}
func main() {
	var b = tmp{
		x: 10,
		y: 20,
	}
	//匿名结构体，没有名字,多用于临时场景
	var a = struct {
		x int
		y int
	}{10, 20} //直接赋了一个初始值
	fmt.Println(a)
	fmt.Println(b)

	var x int
	z := int(10)
	fmt.Println(x, z)

	var p1 person //结构体事例化
	p1.name = "zhoulin"
	p1.age = 19
	p2 := person{"xinjiahui", 18} //结构体事例化
	fmt.Println(p1, p2)
	//	调用构造韩式
	p4 := newPerson("小鹏辅助驾驶lcc", 3)
	fmt.Println(p4)
}

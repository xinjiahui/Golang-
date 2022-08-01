package main

import "fmt"

//结构体遇到的问题
// 1. myInt(100)是个啥
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	//声明一个int32类型的变量x,它的值是10
	//方法1:
	// var x int32
	// x = 10
	//方法二：
	// var x int32 =10
	//方法三：
	// var x = int32(10)
	// 方法4:
	x := int32(10)
	fmt.Println(x)
	//声明一个myInt类型的变量m,它的值是100
	//方法1:
	// var m int32
	// x = 100
	//方法二：
	// var m int32 =100
	//方法三：
	// var m = int32(100)
	//方法四：
	//声明一个int32 类型的
	//问题2:结构体初始化
	type person struct {
		name string
		age  int
	}
	//方法一：
	var p person //声明一个person类型的变量p
	p.name = "xinjiahui"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "test"
	p1.age = 18
	fmt.Println(p1)
	//方法二：
	//键值对初始化
	var p2 = person{
		name: "xian",
		age:  15,
	}
	fmt.Println(p2)
	//值列表初始化：
	var p3 = person{
		"test",
		13,
	}
	fmt.Println(p3)
	p4 := newPerson("辛佳慧", 18)
	fmt.Println(p4)
}

type person struct {
	name string
	age  int
}

//问题3:为什么要有构造函数
func newPerson(name string, age int) *person {
	//别人调用我，我能给他一个person类型的变量
	return &person{
		name: name,
		age:  age,
	}

}

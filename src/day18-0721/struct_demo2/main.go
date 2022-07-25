package main

import "fmt"

// 结构体是值类型
type persion struct {
	name   string
	gender string
}

//go 语言中函数传参数永远是拷贝
func f(x persion) {
	x.gender = "女" //修改的是副本的gender
}
func f2(x *persion) {
	(*x).gender = "女" //根据内存地址找到原变量，修改的就是原来的变量
}

func main() {
	var p persion
	p.name = "haha"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) //男
	f2(&p)
	fmt.Println(p.gender) //女
	//1.结构体指针
	var p2 = new(persion)
	p2.name = "理想"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  //保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) //求p2的内存地址
	// 2.结构体指针2
	//2.1 key-value 初始化
	var p3 = persion{
		name:   "test",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)
	// 2.2使用值列表的形式来初始化，值的顺序更要和结构体的定义时字段的顺序一致
	p4 := persion{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n", p4)
}

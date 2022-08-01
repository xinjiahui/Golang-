package main

import "fmt"

//方法
// 标识符：变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的，就表示对外部包可见（暴露的，公有的）

// dog 这是一个狗的结构体
type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

type persion struct {
	name string
	age  int
}

func newPersion(name string, age int) *persion {
	return &persion{
		name: name,
		age:  age,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写字母表示
func (d dog) wang() {
	fmt.Printf("%s: 汪汪汪～", d.name)
}
func (p *persion) guonian() {
	p.age++
}
func main() {
	d1 := newDog("球球")
	d1.wang()
	p1 := newPersion("terst", 18)
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)

}

package main

import "fmt"

//方法

type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写字母表示
func (d dog) wang() {
	fmt.Printf("%s: 汪汪汪～", d.name)
}
func main() {
	d1 := newDog("球球")
	d1.wang()
}

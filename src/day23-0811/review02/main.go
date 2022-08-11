package main

import "fmt"

type person struct {
	name string
	age  int
}

//方法
//接受者使用对应类型的首个字母小写
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是学好Go语言%s", p.name, str)
}
func main() {
	p1 := person{"zhangjie", 18}
	p1.dream("test")
}

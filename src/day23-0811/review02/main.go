package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
	age  int
}

//方法
//接受者使用对应类型的首个字母小写
//制定了接受者之后，只有接受者这个类型的变量才能调用这个方法
func (p *person) dream(str string) {
	fmt.Printf("%s的梦想是学好Go语言%s", p.name, str)
}

//结构体是值类型
func (p person) guonian() {
	p.age++ //此处的p是p1的副本
}

//指针接受者：
//1.需要修改结构体变量的值时需要使用指针接收者
//2.结构体本身比较大，拷贝的内存开销比较大时也要使用指针接受者
//3.保持一致性：如果有一个方法使用了指针接受者，其他的方法为了统一也要使用指针就收者
func (p *person) guonian1() {
	p.age++
}

//结构体嵌套
type addr struct {
	provice string
	city    string
}
type student struct {
	name    string
	address addr //嵌套别的结构体  名称  类型
}
type student1 struct {
	name string
	addr //匿名嵌套别的结构体，就使用类型名做名称
}

func main() {
	p1 := person{"zhangjie", 18}
	p1.dream("test")
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	p1.guonian1()
	fmt.Println(p1.age)
	type point struct {
		X int `json:"x"` //
		Y int `json:"y"`
	}
	p2 := point{100, 200}
	b, err := json.Marshal(p2) //gong语言中，把错误当成值返回，错误通常作为第二个参数
	if err != nil {
		fmt.Printf("Marshal failed,err:%v\n", err)
	}
	fmt.Println(string(b))
	//反序列化：由字符串--->Go语言中的结构体变量
	str1 := `{"xinjiahui":10,"zhangjie":20}`
	var po2 point //造一个结构体变量，准备接收反序列化的值
	err = json.Unmarshal([]byte(str1), &po2)
	if err != nil {
		fmt.Printf("unmarshal faild,err:%v\n", err)
	}
	fmt.Println(po2)
}

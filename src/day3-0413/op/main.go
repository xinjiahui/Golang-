package main

import "fmt"

func main() {
	var (
		a = 2
		b = 5
	)
	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在=右边赋值 a++ ==> a = a+1
	b--
	fmt.Println(a)
	//关系运算符,go语言是强类型，相同类型的变量才能比较
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
	//逻辑运算符
	// 如果年龄大于18 并且 年龄小于60所
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("上班")
	} else {
		fmt.Println("不用上班")
	}
	//如果年龄小于18岁 或者 年龄大于60岁
	if age < 18 || age > 60 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("上班")
	}
	//not取反，原来为真就为假
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

}

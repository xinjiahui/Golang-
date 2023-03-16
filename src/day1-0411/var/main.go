package main

import "fmt"

//Go 语言中推荐使用驼峰式的命名
//var student_name string
//var studentName string //推荐
//var StudentName string
//声明变量
//var name string
//var age int
//var isOk bool

//批量声明变量
var (
	name string // ""
	age  int    //0
	isOk bool   //false
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	//Go语言中非全局变量声明后必须使用，不使用就编译不过去
	fmt.Print(isOk)             //在终端中输出要打印的内容
	fmt.Println()               //快捷打印空行
	fmt.Println(age)            //打印完指定的内容之后，会在后面加一个换行符
	fmt.Printf("name:%s", name) //%s 占位符，使用name这个变量的值去替换占位符
	fmt.Println()
	//声明变量同时赋值
	var s1 string = "whb"
	//类型推导（根据值判断变量是什么类型）
	var s2 = "xinjiahui"
	fmt.Println(s1)
	fmt.Println(s2)
	//简短变量声明,只能在函数中用
	s3 := "hahaha"
	fmt.Println(s3)

}

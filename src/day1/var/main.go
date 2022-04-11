package main 

import "fmt"

//全局声明变量
var (
	name string // ""
	age int //0
	isOk bool //false
)

func main(){
	name = "理想"
	age = 16
	isOk = true
	//Go语言中非全局变量声明后必须使用
	fmt.Print(isOk)
	fmt.Println()
	fmt.Println(age) //有换行   
	fmt.Printf("name:%s",name) //%s 占位符
	fmt.Println()
	//声明变量同时赋值
	var  s1 string = "whb"
	fmt.Println(s1)
	//简短变量声明,只能在函数中用
	s3 := "hahaha"
	fmt.Println(s3)

}
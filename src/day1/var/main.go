package main 

import "fmt"

//声明变量
var (
	name string // ""
	age int //0
	isOk bool //false
)

func main(){
	name = "理想"
	age = 16
	isOk = true
	//Go语言中声明变量不必使用
	fmt.Print(isOk)
	fmt.Println()
	fmt.Println(age) //有换行   
	fmt.Printf("name:%s",name) //%s 占位符
	fmt.Println()

}
package main

import "fmt"

func main(){
	//age := 45
	if age := 45;age > 35 { //age变量此时只在if条件判断语句中生效
		fmt.Println("人到中年")
	}else if age > 18 {
		fmt.Println("青年")
	}else {
		fmt.Println("好好学习")
	}
}
package main

import "fmt"
// 浮点数
func main(){
    //math.MaxFloat32  
	f1 := 1.234567
	fmt.Printf("%T\n",f1) // 默认Go	语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n",f2) //声明float32类型
    

}
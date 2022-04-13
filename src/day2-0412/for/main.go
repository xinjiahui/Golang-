package main

import "fmt"

//for 循环
func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	//无限循环
	//for {
	//  print.Println("123")
	//}
	// for range 循环
	s := "hello"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
	// 流程控制之跳出for循环
	//当i=5时，跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")
	//当i=5时，跳过此次for循环（不执行for循环内部的打印语句），继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}

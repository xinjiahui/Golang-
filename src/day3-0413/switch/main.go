package main

import "fmt"

// switch 简化大量的判断

func main() {
	var n = 5
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效数字")
	}
	//switch 简化上面的代码
	switch n := 3; n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

}

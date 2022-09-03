package main

import "fmt"

//类型断言

func main() {
	var a interface{} //定一个空接口变量a
	a = 100
	//如何判断a保存值的具体类型
	//类型断言
	//1.x.(T)
	v, ok := a.(int8)
	if ok {
		fmt.Println("猜对了,a是int8", v)
	} else {
		fmt.Println("猜错了,不是int8")
	}

	//2.switch
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case string:
		fmt.Println("string", v2)
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("滚～")
	}
}

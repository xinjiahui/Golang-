package main

import (
	"fmt"
	"strconv"
)

//strconv

func main() {
	//从字符串中解析出对应的数据
	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed,err:", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)
	// 字符串转化成int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)
	//从字符串中解析出bool值
	boolStr := "true"
	booValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", booValue, booValue)
	//从字符串中解析出浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)
	//把数字转换成字符串类型
	i := 97
	//ret2 := string(i)
	ret2 := fmt.Sprintf("%#v", i)
	fmt.Printf("%#v %T\n", ret2, ret2)
	fmt.Println(ret2)
	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v %T\n", ret3, ret3)

}

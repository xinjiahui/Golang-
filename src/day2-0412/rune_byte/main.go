package main

import "fmt"

// byte 和rune类型
//Go语言中为了处理非ASCII码类型的字符。定义了rune类型

func main(){
	s := "Hello辛佳慧신쟈후이"
	//len()求得是byte字节的数量
	n := len(s)
	fmt.Println(n)

	for _, c := range s {
		fmt.Printf("%c\n",c)
	}
	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) //把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))
}
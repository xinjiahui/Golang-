package main

import "fmt"

// byte 和rune类型
//Go语言中为了处理非ASCII码类型的字符。定义了rune类型

func main() {
	s := "Hello辛佳慧신쟈후이"
	s1 := "test"
	fmt.Println(len(s1))
	//len()求得是byte字节的数量
	n := len(s)
	fmt.Println(n)

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) //把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))
	c1 := "红" //string
	c2 := '红' //rune(int32)
	fmt.Printf("c1: %T c2:%T\n", c1, c2)
	fmt.Printf("%d\n", c2)
	c3 := "H" //string
	c4 := 'H' //int32
	fmt.Printf("c3: %T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	//类型转换
	m := 10 //int
	var f float64
	f = float64(m)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}

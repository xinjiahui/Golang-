package main

import "fmt"

// fmt

func main() {
	fmt.Print("nihao")
	fmt.Print("hahah")
	fmt.Println("xhj")
	fmt.Println("hello")
	// Printf("格式化字符串"，值)
	// %T : 查看类型
	// %d : 十进制数
	// %b : 二进制数
	// %o : 八进制数
	// %x : 十六进制数
	// %c : 字符
	// %s : 字符串
	// %p : 指针
	// %v : 值
	// %f : 浮点数
	// %t : 布尔值
	var m1 = make(map[string]int, 1)
	m1["李想"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)
	PrintBaifenbi(90)
	// fmt.Printf("%s\n",100)
	// 整数->字符
	fmt.Printf("%q\n", 65)
	// 浮点数和复数
	fmt.Printf("%b\n", 3.14159265354697)
	fmt.Printf("%q\n", "李想有理想")
	// 获取用户输入
	var s string
	fmt.Scan(&s) //Scanf Scanln
	fmt.Println("用户输入的内容是： ", s)
}
func PrintBaifenbi(num int) {
	fmt.Printf("%d%%\n", num)
}

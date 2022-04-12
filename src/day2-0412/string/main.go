package main

//Go语言中字符串是用双引号包裹的！
//Go语言中单引号包裹的的字符！

// 字符串
import "fmt"
import "strings"

func main(){
	path  := "\"D:\\Go\\src\\code\\day01\""
	fmt.Println(path)

	s := "I am ok"
	fmt.Println(s)
	//定义多行字符串
    s2 := `
	世情薄
	人情恶
	雨送黄昏花易落
	`
	fmt.Println(s2)
	s3 := `D\Go\src\code\dat01`
	fmt.Println(s3)
	fmt.Println(len(s3))
	//字符串拼接
	name := "jiahui"
	world := "haoren"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s",name,world)
	fmt.Println(ss1)
	//分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss, "jiahui"))

	// 前缀
	fmt.Println(strings.HasPrefix(ss,"jiahui"))
	fmt.Println(strings.HasSuffix(ss,"hui"))

	// 拼接
	fmt.Println(strings.Join(ret,"+"))

}
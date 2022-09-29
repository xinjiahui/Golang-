package main

import (
	"flag"
	"fmt"
	"time"
)

//flag	获取命令行参数
func main() {
	//创建一个标志位参数
	name := flag.String("name", "辛佳慧", "请输入名字")
	age := flag.Int("age", 18, "请输入真实年龄")
	married := flag.Bool("married", false, "结婚了吗？")
	ctime := flag.Duration("ct", time.Second, "结婚多久了")
	//使用flag
	flag.Parse() //先解析在使用
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*ctime)
	fmt.Println(flag.Args())  //返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用命令行参数的个数
}

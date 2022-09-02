package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入时，如果有空格

func useScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是：%s\n", s)
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin) //从哪里读文件，fileObj就是从打开的文件中读，Stdin是从标准输入读。相同，写也是可以写入文件，也可以写入终端 是一个接口（interface{}）
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
}

func main() {
	useScan()
	useBufio()
	//写日志
	fmt.Fprintln(os.Stdout, "这是一条日志记录！")
	fileObj, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file faild,err:%v", err)
	}
	fmt.Fprintln(fileObj, "这还是一条日志记录！")
}

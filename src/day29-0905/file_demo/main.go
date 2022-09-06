package main

import (
	"fmt"
	"os"
)

//文件对象的类型
//获取文件对象的详细信息
func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
	}
	fmt.Println("fdsfsfdsgsdgsg")
	//文件对象的类型
	fmt.Printf("%T\n", fileObj)
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return
	}
	fmt.Printf("文件大小是：%dB\n", fileInfo.Size())
	fmt.Printf("文件名字：%s\n", fileInfo.Name())
}

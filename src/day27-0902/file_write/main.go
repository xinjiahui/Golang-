package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容
func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) //逻辑或操作，有一个为1，就为1
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	//write
	fileObj.Write([]byte("you dian yi huo\n"))
	//writestring
	fileObj.WriteString("哈哈哈哈")

	fileObj.Close()
}
func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) //逻辑或操作，有一个为1，就为1
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("你好呀～\n") //写到缓存中
	wr.Flush()               //将缓存中的内容写入文件
}
func writeDemo3() {
	str := "等一朵花开～"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file failed,err:%v", err)
		return
	}

}
func main() {
	writeDemo1()
	writeDemo2()
	writeDemo3()
}

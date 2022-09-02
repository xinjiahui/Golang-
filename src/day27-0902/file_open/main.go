package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("Open file faild,err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	//读文件
	var tmp [128]byte //指定读的长度
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file faild,err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

//利用bufio这个包来读取文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed.err:%v", err)
		return
	}
	//记得关闭对象
	defer fileObj.Close()
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v", err)
			return
		}
		fmt.Print(line)
	}

}
func readFromFileByIouil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed,err:%v", err)
		return
	}
	fmt.Println(string(ret))

}
func main() {
	//readFromFile1()
	//readFromFilebyBufio()
	readFromFileByIouil()
}

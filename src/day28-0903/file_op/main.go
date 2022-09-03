package main

import (
	"fmt"
	"io"
	"os"
)

func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()

}
func f11() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./xxx.go")
	defer fileObj.Close() //写到这里不可以，当err有值的时候，fileObj是nil。nil不能调用close方法,尽量不要这样
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}

}
func f2() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	//defer fileObj.Close()

	// 因为没有办法直接在文件中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./test.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed,err:%v\n", err)
	}
	defer tmpFile.Close()
	//读源取文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed,err:%v", err)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])
	// 在写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)
	//紧接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	//源文件后续的也写入了临时文件
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./test.tmp", "./test.txt")
}
func main() {
	//f1()
	//f11()
	f2()
}

package main

import (
	"fmt"
	"path"
	"runtime"
)

//runtime.Caller()
func f1() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() failed \n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name() //函数名字
	fmt.Println(funcName)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
	fmt.Println(ok)

}
func main() {
	f1()
}

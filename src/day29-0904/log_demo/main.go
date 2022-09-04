package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//支持往不同的地方输出日志
//日志分级别
// Debug
// Info
// Trace
// Warning
// Error
// Fatal

// log demo

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed.err:%v\n", err)
		return
	}
	log.SetOutput(fileObj)
	for {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 3)
	}
}

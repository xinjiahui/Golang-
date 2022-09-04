package main

import (
	"Golang-/src/day29-0904/mylogger"
	"time"
)

//测试我们自己写的日志

func main() {
	//log := mylogger.NewLog("Error")
	log := mylogger.NewFileLogger("Error", "./", "xinjiahui.log", 10*1024*1024)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条warning日志")
		id := 10010
		name := "辛佳慧"
		log.Error("这是一条error日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条fatal日志")
		time.Sleep(time.Second)
	}
}

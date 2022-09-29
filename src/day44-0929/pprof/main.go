package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode() {
	var c chan int //nil
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan,value:%v\n", v)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUPprof bool //是否开启CPUprofile的标识位
	var isMemPprof bool //是否开启内存profile的标识位
	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()
	if isCPUPprof {
		file1, err := os.Create("./cpu.pprof") //在当前路径下创建一个cpu.pprof文件
		if err != nil {
			fmt.Printf("create cpu pprof failed,err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file1) //往文件中记录CPU profile信息
		defer func() {
			pprof.StopCPUProfile()
			file1.Close()
		}()

	}
	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof fiaid,err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

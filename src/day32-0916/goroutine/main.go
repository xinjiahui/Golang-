package main

import (
	"fmt"
	"time"
)

//goroutine

func hello(i int) {
	fmt.Println("hello", i)
}

//程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		//go hello(i) //开启一个单独的goroutine 去执行hello函数 （任务）
		//匿名函数
		go func(i int) {
			fmt.Println(i) //用的是函数参数的那个i，不是外面的i
		}(i)
	}
	fmt.Println("main")
	//main函数结束了由main函数启动的goroutine也都结束了
	time.Sleep(time.Second)

}

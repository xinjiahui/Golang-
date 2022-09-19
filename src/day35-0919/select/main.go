package main

import "fmt"

//select 多路通道操作

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		//如果都满足，走哪个分支是随机的
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Print("hahaha")
		}
	}
}

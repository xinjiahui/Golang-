package main

import (
	"fmt"
	"math/rand"
	"time"
)

//channel
//往通道中发送值
func sendNum(ch chan<- int) {
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(5 * time.Second)
	}

}

func main() {
	ch := make(chan int, 1)
	go sendNum(ch)
	for {
		x, ok := <-ch //等4s钟执行
		fmt.Println(x, ok)
		time.Sleep(time.Second)
	}

}

package main

import (
	"fmt"
	"sync"
)

//var a []int
//需要指定通道中元素的类型
var b chan int
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)
	b = make(chan int) //不带缓冲区通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10
	fmt.Println("发送到通道b中了...")
	wg.Wait()

}
func BufChannel() {
	fmt.Println(b)         //nil
	b = make(chan int, 16) //带缓冲区通道的初始化
	b <- 30
	fmt.Println("30发送到通道b中了...")
	b <- 40
	fmt.Println("40发送到通道b中了...")
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)
}
func main() {
	noBufChannel()
	BufChannel()
}

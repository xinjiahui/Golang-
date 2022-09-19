package main

import "fmt"

//关闭通道
var ch1 chan<- int

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	//for x := range ch1 {
	//	fmt.Println(x)
	//}
	x, ok := <-ch1
	fmt.Println(x, ok) //10
	y, ok := <-ch1
	fmt.Println(y, ok) //20
	z, ok := <-ch1
	fmt.Println(z, ok) //0
}

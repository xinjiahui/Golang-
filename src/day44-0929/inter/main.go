package main

import "fmt"

//如何判断一个链表有没有闭环
//x走1步，y走2步,如果x,y在某一时刻相遇了，有闭环
//type a struct {
//	val  int
//	next *a
//}

// 台阶的问题：有n阶楼梯，一次可以走1个台阶，一次可以走2个台阶，问有多少种走法
func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)
}

func main() {
	fmt.Println(f(10))
}

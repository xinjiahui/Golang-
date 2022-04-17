package main

import (
	"fmt"
	"sort"
)

// 		切片的练习题
func main() {
	var a = make([]string, 5, 10) // 创建长度为5容量为10的切片
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Println(cap(a))
	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:]) // 对切片进行排序
	fmt.Println(a1)
}

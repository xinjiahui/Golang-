package main

import "fmt"

//数组练习题
//1.求数组[1,3,5,7,8]所有元素和
//2.找出数组中和为指定值的两个元素的下标。比如从数组【1，3，5，7，8】中找出和为8的两个元素的下标为（0，3）和（1，2）
func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a1 {
		sum = v + sum
	}
	fmt.Println(sum)
	// 题2
	//定义2个for循环
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}

}

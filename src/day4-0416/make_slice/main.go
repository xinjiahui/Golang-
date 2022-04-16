package main

import "fmt"

//make  函数创造切片
// 切片就是一个框，框住了一块连续的内存。属于引用类型，真正的数据都是保存在底层数组里的
// 判断一个切片是否为空，要用len（s）== 0 来判断
func main() {
	s1 := make([]int, 5, 10) //长度为5.容量为10
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := make([]int, 0, 10) //长度为0.容量为10
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))
	//切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3 //s3和
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	//切片的遍历
	//1.索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	for i, v := range s3 {
		fmt.Println(i, v)
	}
}

package main

import "fmt"

//append()为切片 追加元素
func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s1[3] = "广州" //错误的写法，会导致编译错误：索引越界
	//fmt.Println(s1)
	// 调用append函数必须用原来的切片变量接收返回值
	s1 = append(s1, "广州") //append追加元素，原来的底层数组放不下的时候，go底层就会把底层数组换一个
	fmt.Println(s1)
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) // ...	表示拆开
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// append删除切片中的某个元素
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s2 := a1[:]
	// 删除索引为1的那个3
	s2 = append(s2[0:1], s2[2:]...)
	fmt.Println(s2)
	fmt.Println(a1)
}

package main

import "fmt"
// 数组
// 存放元素的容器
//必须指定存放元素的类型和容量（长度）
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool //[true false true]
	var a2 [4]bool //[true true false false]
	fmt.Printf("a1: %T a2:%T", a1, a2)
	// 数组的初始化
	// 如果不初始化，默认元素都是0值。（布尔值：false，整型和浮点型都是0，字符串：“”）
	fmt.Println(a1, a2)
	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2.初始化方式2
	// a10 := [9]int{0,1,2,3,4,5,6,7,8}
	//根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	// 3.初始化方式3:根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	// 1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2.for range 遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}
	//多维数组，只有最外层可以使用...
	// [[1 2][3 4] [5 6]]
	var a11 [3][2]int //三个元素，每个元素是长度为2的int类型数组
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)
	// 多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}

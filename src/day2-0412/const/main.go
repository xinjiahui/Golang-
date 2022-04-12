package main

import "fmt"
//常量
//定义了常量之后不能修改
//在程序运行期间不会改变的量
const pi = 3.1415926
// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2 
	n3
)

// iota
// const 中每新增一行常量声明将均使iota计数一次。使用iota能简化定义，在定义枚举时很有用。

const (
	a1 = iota // 0
	a2  //1
	a3 //2
)

const (
	b1 = iota // 0
 	b2 = iota //1
	_          //2 
	b3 = iota //3
)


//定义数量级

const(
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota) 
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)	
	PB = 1 << (10 * iota)
)
func main(){
    fmt.Println(a1,a2,a3)

}
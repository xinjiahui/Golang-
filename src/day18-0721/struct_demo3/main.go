package main

//内存对齐:CPU 访问内存时，并不是逐个字节访问，而是以字长（word size）为单位访问
//比如 32 位的 CPU ，字长为 4 字节，那么 CPU 访问内存的单位也是 4 字节
//字长定义在同一时间中处理二进制数的位数叫字长。通常称处理字长为8位数据的CPU叫8位CPU，32位CPU就是在同一时间内处理字长为32位（4字节）的二进制数据。
//二进制的每一个0或1是组成二进制的最小单位，称为位（bit）。常用的字长为8位、16位、32位和64位。字长为8位的编码称为字节，是计算机中的基本编码单位。

import (
	"fmt"
	"unsafe"
)

//结构体占用一段连续的内存空间
type x struct {
	a int8 //8bit=1byte
	b string
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: "哈哈",
		c: int8(30),
	}
	fmt.Printf("%p\n", &m.a)
	fmt.Printf("%p\n", &m.b)
	fmt.Printf("%p\n", &m.c)
	fmt.Println(unsafe.Sizeof(x{}))
	fmt.Println(unsafe.Alignof(x{}))
}

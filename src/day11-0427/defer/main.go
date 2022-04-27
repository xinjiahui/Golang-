package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x 不是返回值
	}()
	return x // 1.返回值复制 2.defer 3.真正RET指令
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值 = x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是x
	}()
	return x // 返回值 = y = x =5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //返回值 = x
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}

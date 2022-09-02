package main

import (
	test "Golang-/src/day27-0902/package/hahha-calc" //别名
	"fmt"
)

func init() {
	fmt.Println("自动执行")
}

func main() {
	ret := test.Add(10, 20)
	fmt.Println(ret)
}

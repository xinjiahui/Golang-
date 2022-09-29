package main

import (
	"Golang-/src/day43-0928/split_string"
	"fmt"
)

func main() {
	ret := split_string.Split("backfsbsf", "b")
	fmt.Printf("%#v\n", ret)
	ret2 := split_string.Split("bbb", "b")
	fmt.Printf("%#v\n", ret2)
	ret3 := split_string.Split("fesfsf", "b")
	fmt.Printf("%#v\n", ret3)
}

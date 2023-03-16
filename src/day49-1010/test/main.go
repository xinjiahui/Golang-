package main

import "fmt"

func main() {
	var s int
	m := string(s)
	if m == "" {
		fmt.Println("test")
	}
	fmt.Println(m)
}

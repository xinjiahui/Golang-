package main

import "fmt"

func yuanshuai(name string) {
	fmt.Println("Hello", name)
}

func lixiang(f func(string), name string) {
	f(name)
}
func main() {
	lixiang(yuanshuai, "xinjiahui")
}

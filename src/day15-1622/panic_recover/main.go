package main

import "fmt"

//panic 和recover
func f1() {
	defer func() {
		err := recover()
		fmt.Println("我好难过呀...你们都是坏人...")
		fmt.Println(err)
	}()
	panic("盈盈一水，或许我掬不起最清澈的;姹紫嫣红，或许我携不来最芬芳的;皓月当空，或许我等不来最明朗的。但是我可以轻触，可以浅嗅，可以在无数的残缺中体会圆满...")
	fmt.Println("f1")
}
func f2() {
	fmt.Println("f2")
}

func main() {
	f1()
	f2()

}

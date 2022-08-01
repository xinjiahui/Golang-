package main

import "fmt"

/*
	函数版学生管理系统
	写一个系统能够查看、新增，删除学生
*/

type student struct {
	id   int64
	name string
}

func main() {
	//1.打印菜单
	fmt.Println("欢迎管理学生管理系统")
	fmt.Println(`1.查看所有学生
		 2.新增学生
		 3.删除学生
	`)
	fmt.Println("请输入你要干啥：")
	var choice int
	fmt.Scanln(&choice)
	fmt.Printf("你选择了%d这个选项!\n", choice)
	//2.等待用户选择要做什么
	//3.执行对应的函数
}

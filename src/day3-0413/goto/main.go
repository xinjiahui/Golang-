package main

import "fmt"

func main() {
	//正常退出for循环，常用的
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'B' {
				flag = true
				break
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break //跳出for循环（外层的for循环）
		}
	}
	//goto跳出for循环，不建议使用
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'B' {
				//设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%c\n", i, j)
		}

	}
breakTag: //标签
	fmt.Println("结束for循环")
}

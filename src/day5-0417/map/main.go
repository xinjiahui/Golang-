package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间动态扩容
	m1["理想"] = 18
	m1["jiwuling"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["test"])
	value, ok := m1["test"]
	if !ok {
		fmt.Println("查无此Key")
	} else {
		fmt.Println(value)
	}
	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	//只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	//删除
	delete(m1, "jiwuling")
	fmt.Println(m1)

}

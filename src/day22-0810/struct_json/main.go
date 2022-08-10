package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json

//1.序列化：把Go语言中的结构体变量--->json格式的字符串
//2.反序列化：json格式的字符串---> Go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"` //Name不大写在外面调用不到
	Age  int    `json:"age"`                       //json中的字段是小写的
}

func main() {
	p1 := person{
		Name: "张杰",
		Age:  40,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal faild,err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	//反序列化
	str := `{"name":"张杰","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在json.unmarshal内部修p2的值
	fmt.Printf("%#v\n", p2)
}

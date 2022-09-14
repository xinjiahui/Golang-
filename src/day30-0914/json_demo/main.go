package main

import (
	"encoding/json"
	"fmt"
)

// 反射
//json

type persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"xinjiahui","age":9000}`
	var p persion
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
}

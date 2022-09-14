package main

import (
	"fmt"
	"reflect"
)

//反射
//反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分，在运行程序时，程序无法获取自身的信息。json.Unmarshal()
//性能低下：基于反射实现的代码运行速度慢一到两个数量级
type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind:%v\n\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() //值的类型种类
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func main() {
	var a float32 = 3.14
	reflectType(a) //type:float32
	var b int64 = 100
	reflectType(b)
	var c = Cat{}
	reflectType(c)
	reflectValue(a)
	// 设置值
	reflectSetValue1(&b)
	fmt.Println(b)
	reflectSetValue2(&b)
	fmt.Println(b)

}

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini 配置文件解析器

//MysqlConfig Mysql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `int:"port"`
	Username string `int:"username"`
	Password string `int:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//0.参数的校验
	//0.1 传进来的data参数必须是指针类型（因为要在函数中对其赋值）
	t := reflect.TypeOf(data)
	//fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer") //创建错误
		return
	}
	//0.2 传进来的data参数必须是结构体类型指针（因为配置文件中有各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer") //创建错误
		return
	}
	//1.读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b) //将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("%#v\n", lineSlice)
	//2.一行一行得度数据
	var structName string
	for idx, line := range lineSlice {
		//去掉字符串首尾空格
		line := strings.TrimSpace(line)
		//如果是空行，就跳过
		if len(line) == 0 {
			continue
		}
		//2.1如果是注释，就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//2.2 如果是[ 开头的，就表示节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d suntax error", idx+1)
				return
			}
			//把这一行收尾的[]去掉，取到中间的内容把首位的空格去掉拿掉内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//根据字符串sectionName去data里面根据反射找到对应结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了嵌套的结构体，把字段名字记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			//2.3 如果不是[	开头，就是=分割的键值对
			//1.以等号分割这一行，等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			//fmt.Printf("key:%s----\n", key)
			value := strings.TrimSpace(line[index+1:])
			//2.根据structName，去data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			//fmt.Println(sValue)
			sType := sValue.Type() //拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			//3.遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) //tag信息是存储在类型信息中的
				fileType = filed
				//fmt.Println(fileType.Tag.Get("ini"))
				if filed.Tag.Get("ini") == key {
					//	找到了对应的字段
					fieldName = filed.Name
					break
				}
			}
			//4.如果key= tag,给这个字段赋值
			//4.1根据fieldName去取出这个字段

			if len(fieldName) == 0 {
				//在结构体中找不到对应的字符
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			//4.2 对其赋值
			fmt.Printf("%v: %v\n", fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}

		}

	}

	return

}
func main() {
	var cfg Config
	err := loadIni("./config.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}

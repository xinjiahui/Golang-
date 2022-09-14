package main

import "fmt"

//ini 配置文件解析器

//MysqlConfig Mysql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `int:"port"`
	Username string `int:"username"`
	Password string `int:"password"`
}

func loadIni(v interface{}) {

}
func main() {
	var mc MysqlConfig
	loadIni(&mc)
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)

}

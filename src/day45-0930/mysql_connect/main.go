package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

// go 连接mysql示例

//下载驱动go get -u github.com/go-sql-driver/mysql
// 原声支持连接池，是并发安全的，这个标准库没有具体的实现，只是列出了一些需要的第三方库实现的具体内容
func main() {
	//数据库信息
	dsn := "root:insecure-root-password@tcp(10.5.42.101:31912)/xinjiahui_gotest"
	//连接数据库
	db, err := sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {                   //dsn格式不正确会报错
		fmt.Printf("dsn:%s invalid,err:%v", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed ,err:%v", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")
}

package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // init()
)

var db *sqlx.DB

func initDB() (err error) {
	//数据库信息
	dsn := "root:SenseStudy2018@tcp(127.0.0.1:3306)/xinjiahui_gotest"
	//连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	//设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

//SQL注入示例

func sqlInjectDemo(name string) {
	//自己拼接SQL语句的字符串
	sqlStr := fmt.Sprintf("select id,name,age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed ,err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user%#v\n", u)
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
	}
	//SQL注入的几种示例
	sqlInjectDemo("jason") //正常
	sqlInjectDemo("xxx'or 1=1 #")
}

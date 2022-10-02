package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
	"github.com/jmoiron/sqlx"
)

//sqlx demo
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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	sqlStr1 := `select id,name,age from user where id=1`
	var u user
	err = db.Get(&u, sqlStr1)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
	}
	fmt.Printf("u:%#v\n", u)
	var userList []user
	sqlStr2 := `select id,name,age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
	}
	fmt.Printf("userList:%#v\n", userList)

}

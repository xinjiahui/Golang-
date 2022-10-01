package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

// go 连接mysql示例
var db *sql.DB //是一个连接池对象
func initDB() (err error) {
	//数据库信息
	dsn := "root:SenseStudy2018@tcp(127.0.0.1:3306)/xinjiahui_gotest"
	//连接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {                  //dsn格式不正确会报错
		return
	}
	err = db.Ping()
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
	id   int
	name string
	age  int
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v\n", err)
		return
	}
	//执行多个SQL操作
	sqlStr1 := `update user set age=age-2 where id = 1`
	sqlStr2 := `update user set age=age-2 where id = 8`
	//执行SQL
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行SQL1出错了,要回滚")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行SQL2出错了,要回滚")
		return
	}
	//上面两步都执行成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("提交出错了,要回滚")
		return
	}
	fmt.Println("事务执行成功")
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功")
	transactionDemo()
}

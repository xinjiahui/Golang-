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

//预处理方式插入多条数据

func prepareInsert() {
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr) //把SQL语句先发给mysql预处理一下
	if err != nil {
		fmt.Printf("prepare failed ,err:%v\n", err)
		return
	}
	defer stmt.Close()
	//后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"皱皱":    13,
		"王野":    15,
		"test":  14,
		"ligai": 15,
	}
	for k, v := range m {
		stmt.Exec(k, v) //后续只需要传值
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功")
	prepareInsert()
}

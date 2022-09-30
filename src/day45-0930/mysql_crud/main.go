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

//查询单个记录
func queryOne(id int) {
	//1.写查询单条记录的sql语句
	sqlStr := "select id,name,age from user where id=?"
	//2.执行
	rowObj := db.QueryRow(sqlStr, id) //从连接池里拿一个连接出来去查询单条记录
	//3.拿到结果
	var u1 user
	rowObj.Scan(&u1.id, &u1.name, &u1.age) //Scan 释放数据库连接到连接池
	//4.打印结果
	fmt.Printf("u1:%#v\n", u1)
}

//查询多条
func queryMore(n int) {
	//1.SQL 语句
	sqlStr := `select id,name,age from user where id > ?`
	//2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed,err:%v\n", sqlStr, err)
		return
	}
	//一定要关闭rows
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan  failed,err:%v\n", err)
			return
		}
		fmt.Printf("u:%#v\n", u)

	}

}

func insert() {
	//写SQL语句
	sqlStr := `insert into user(name,age) values("jason",40)`
	//执行
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	//如果是插入数据的操作，能够拿到插入数据的ID值
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastid failed,err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get idfailed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据", n)
}

func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get idfailed,err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据", n)
}

//下载驱动go get -u github.com/go-sql-driver/mysql
//原声支持连接池，是并发安全的，这个标准库没有具体的实现，只是列出了一些需要的第三方库实现的具体内容
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功")
	queryOne(1)
	queryMore(0)
	insert()
	updateRow(900, 4)
	deleteRow(3)
}

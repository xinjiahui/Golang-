### Mysql   
- 数据库：
    - 常见的数据库：SQLlite,Mysql,SQLServer,postgreSQL,Oracle
    - mysql主流的关系性数据库，类似的还有postgreSQL
    - 关系型数据库
- MySQL 知识点
    - SQL语句
    DDL:操作数据库的
    DML:表的增删改查
    DCL:用户及权限
    - 存储引擎
    mysql支持插件式的存储引擎。
    常见的存储引擎：MyISAM和InnoDB
        - My ISAM：
            1.查询速度快
            2.只支持表锁
            3.不支持事物
        - InnoDB   
            1.整体速度快
            2.支持表锁和行锁
            3.支持事务 
- 事务
    - 把多个SQL操作当成一个整体
- 事务的特点：
    - ACID：
        1.原子性：
        2.一致性：
        3.隔离性：
            隔离级别：
        4.持久性：
- 索引：
    - B树和B+树
    - 索引的类型
    - 索引的命中
- 分库分表
- SQL慢查询优化
- SQL注入
- MySQL主从（binlog）
- 读写分离

### Go操作数据库

- 下载驱动
```
go get -u github.com/go-sql-driver/mysql 
```
- 原声支持连接池，是并发安全的，这个标准库没有具体的实现，只是列出了一些需要的第三方库实现的具体内容
- go get 包的路径就是下载第三方的依赖，将第三方的依赖默认保存在 $GOPATH/src/
## package
### 包的路径
- 从GOPATH/src 后面的路径开始写，路径分隔符是 / 
- 想被别的包调用的标识符都要首字母大写！
- 单行导入和多行导入
- 导入包的时候可以指定别名
- 导入包不想使用包内部的标识符需要使用匿名导入
- 每个包导入的时候会自动执行一个init（）的函数，它没有参数，没有返回值，也不能手动调用
- 多个包中都定义了init（）函数，，则它们的执行顺序如下图：
image.png

## 日志模块

- 


## 日志库
- 自己写一个日志库
- 接口：用处？日志可以输出到终端，也可以输出到文件，输出到kafka
- 需求：
    - 1.可以往不同的输出位置记录日志
    - 2.日志分为5种级别
        - 1.

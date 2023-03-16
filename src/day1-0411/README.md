### 变量和常量

- 标识符与关键字
    - 标识符
        - 在编程语言中标识符就是程序员定义的具有特殊意义的词，比如变量名，常量名，函数名等等。Go语言中标识符由字母数字和_(下划线)组成，并且只能以字母和 _ 开头。 abc,_,_123,a123
    - 关键字
        - 关键字是指编程语言中预先定义好的具有特殊含义的标识符，关键字和保留字都不建议用作变量名
        - Go 语言中有25个关键字
        ```
        break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var
        ```
    - 变量
        - Go 语言中的变量，必须先声明在使用
        - 声明变量
            ```
            var s1 string
            ```
             声明一个保存字符串类型的变量。
        - Go 语言中非全局变量声明了必须使用
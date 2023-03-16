### Go module

- go module 是Go 1.11版本之后，官方推出的版本管理工具，并且从Go1.13版本开始，go module将是Go语言默认的依赖管理工具。

#### Go111MODULE
- 要启用go module 支持，首先要设置环境变量GO111MODULE,通常它可以开启或关闭模块支持，它有3个可选值：off，on,auto。默认值auto
    - 1. GO111MODULE=off 禁用模块支持，编译时会从GOPATH和vendor文件夹中查找。
    - 2. GO111MODULE=on 启用模块支持，编译时会忽略GOPATH和verdor文件夹，只根据go.mod 下载依赖。//在pkg目录下
    - 3. GO111MODULE=auto,当项目在$GOPATH/src外且项目跟目录有go.mod文件夹时，开启模块支持。

```
export GO111MODULE=on
```
- 简单来说，设置GO1111MODULE=on之后就可以用go module了，以后就没必要在GOPATH中创建项目了。并且还能够很好的管理项目的第三方包信息。
- 使用go module管理依赖后会在项目跟目录下生成2个文件，go.mod和go.sum


#### GOPROXY

- Go1.11之后设置GOPROXY命令为：
```
export GOPROXY=https://goproxy.cn
```
- Go1.13版本之后GOPROXY默认值为 https://proxy.golang.org,在国内是无法访问的，所以使用https://goproxy.cn

```
go env -w GOPROXY=https://goproxy.cn,direct
```

#### go mod 命令

- go mod download #下载依赖的module到本地cache(默认为$GOPATH/pkg/mod)
- go mod edit #编辑go.mod文件
- go mod graph # 打印模块依赖图
- go mod init #初始化当前文件夹，创建go.mod文件
- go mod tidy # 增加缺少的module，删除无用的module
- go mod vendor # 将依赖复制到vendor下
- go mod verify # 校验依赖
- go mod why # 解释为什么需要依赖

#### go.mod
- go.mod记录了项目所有的依赖信息，其结构大致如下：
```
module github.com/xinjiahui/test/test  
go 1.18.4

require(
    github.com/go-sql-driver/mysql v1.4.1
    google.golang.org/appengine v1.6.1 //indriect
)
replace(
    golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```
- 1. module用来定义包名
- 2. require用来定义依赖包及版本
- 3. indirect 表示间接引用
- 4. replace 在国内访问 golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库



```
1. export GO111MODULE=on
2. go mod init [包名] //初始化项目
3. go mod tidy  检查依赖
4. export GOPROXY=https://goproxy.cn 
5. go get
```



### Context
- Go 1.7加入了一个新的标准库context，它定义了Context类型，专门用来简化对于处理单个请求多个goroutine之间与请求域的数据， 取消信号，截止时间等相关操作这些操作可能涉及多个API调用。
- 对服务器传入的请求应该创建上下文，而对服务器的传出应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel，WithDeadline,WithTimeout,或WithValue 创建的派生上下文，当一个上下文被取消时，它派生的所有上下文也被取消。


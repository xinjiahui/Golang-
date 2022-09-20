## 内容回顾

### 并发之goroutine
- 并发和并行的区别
- goroutine启动
    - 将要并发执行的任务包装成一个函数，调用函数的时候前面加上go关键字，就能够开启一个goroutine去执行该函数
- goroutine结束
    - groutine对应的函数执行完，该goroutine就结束了。
    - main函数启动的时候会自动创建一个goroutine 去执行main函数
    - main函数结束了，那么程序就结束了，由该程序启动的所有其他goroutine也都结束了
- sync.WaitGroup]
```
var wg sync.WaitGroup
```
    - wg.Add(1) //计数器 +1
    - wg.Done() // goroutine内部执行 计数器-1
    - wg.Wait() //等
- goroutine的本质       
    - goroutine的调度模型GMP
    - 把m个goroutine分配给n个操作系统线程
- gorotine操作系统（OS）线程的区别
    - goroutine是用户态的线程，比内核态的线程更轻量级一点，初始时只占用2KB的栈空间。可以轻松开启数十万的goroutine也不会崩内存
- runtime.GOMAXPROCS
    - Go 1.5 之后默认就是操作系统的逻辑核心数，默认跑慢CPU
    ```
    runtime.GOMAXPROCS(1) //只占用一个核心
    ```
- work pool 模式
    - 开启一定数量的goroutine去干活

### channel
- 为什么需要channel？
    - 通过channel实现多个goroutine之间的通信
    - CSP,通过通信来共享内存
- channel 是一种类型，是一种引用类型 make函数初始化之后才能使用（slice,map，channel）
    - channel 的声明
    ```
    var ch chan 元素类型
    ```
    - channel的初始化
    ```
    ch = make(chan 元素类型，[缓冲区大小])
    ```
- channel的操作
    - 发送
    ```
    ch <-100
    ```
    - 接收
    ```
    x := <-ch
    ```
    - 关闭
    ```
    close(ch)
    ```
- 带缓冲区的通道和无缓冲区的通道：
    - 关闭已经关闭的通道，会painc
- 单向通道
    - 通常是用作函数的参数，只读通道<-chan  int 只写通道 chan <- int 
    
- select 
    - 同一时刻有多个通道要操作的场景。Go内置了Select关键字，可以同时响应多个通道的操作。
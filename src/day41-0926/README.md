### UDP协议。
- 一种无连接的传输层协议，不需要建立连接就能直接进行数据发送和接收，属于不可靠的，没有时序的通信。但是UDP协议的实时性比较好，通常用于视屏直播相关领域。
### 重要的！！！
- 注释
- 日志
- 单元测试

### 锁
- sync.Mutex //是一个结构体，是值类型，给函数传参数的时候要传指针
- 两个方法：
```
var lock sync.Mutex
lock.Lock()//加锁
lock.Unlock() //解锁
```
- 为什么要用锁：
    - 防止同一时刻多个goroutine操作同一个资源

### 读写互斥锁
- 应用场景：
    - 适用于度多写少的场景，才能够提高程序的执行效率
- 特点：
    - 读的goroutine来了获取的是读锁，后续的goroutine能读不能写
    - 写的goroutine来了获取的是写锁，后续的goroutine无论是读还是写都要等待获取锁
- 使用：
```
var rwlock sync.RWMutex
rwlock.RLock() //获取读锁
rwlock.RUnlock() // 获取读写
rwlock.Lock() // 获取写锁
rwlock.Unlock() //释放写锁
```
### 等待组
```
sync.Waitgroup
```
- 本质上也是一个结构体，值类型，给函数传参数的时候要传指针。用来等groutine执行完在继续
- 使用
```
var wg sync.WaitGroup
wg.Add(1) //起几个goroutine就加几个计数
wg.Done() //在goroutine 对应的函数中，函数要结束的时候表示goroutine完成，计数器-1
wg.Wait() //阻塞，等待所有的goroutine都结束
```
###  Sync.Once
- 使用场景
     - 某些函数只需要执行一次的时候，就可以使用`sync.once`,比如blog加载图片
     ```
     var once sync.Once
     once.Do() //接收一个没有参数，没有返回值的函数。如有需要可以使用闭包
     ```
### sync.Map
- 使用场景
    - 并发操作一个map的时候，内置的map不是并发安全的。
- 使用
    - 是一个开箱即用(不需要make初始化)的并发安全的map
    ```
    var syncMap sync.Map
    syncMap[key] = value // 原生map
    syncMap.Store(key,value)
    syncMap.Load(key)
    syncMap.LoadOrStore
    syncMap.Delete()
    syncMap.Range()
    ```
### 原子操作
- Go 语言内置了一些针对内置的基本数据类型的一些并发安全的操作
```
var i int64 = 10
atomic.AddInt64(&i,1)
```

### 网络编程
- 互联网的协议
    - OSI七层模型：应（SMTP,HTTP,FTP），表，会，传(TCP,UDP)，网（IP协议,MAC地址），数（以太网协议），物(双绞线)
- TCP server client
- 粘包
- UDP server client





### goroutine 什么时候结束
- goroutine对应的函数结束了，goroutine结束了。
- main 函数执行完了，由main函数创建的那些goroutine都结束了。

- 保证每次执行的结果不一样
```
rand.Seed(time.Now().UnixNano()) //随机数的种子
```


### 理论的知识

- goroutine 与线程
#### GMP是Go语言的调度模型
    - GMP是Go语言运行时层面的实现，是go语言自己实现的一套调度系统，区别于操作系统调度os线程。
        - G很好理解，就是gorontine的，里面除了存放本goroutine信息外，还有与所在P的绑定信息。
        - M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟，M与内核线程一般是一一映射的关系，一个groutine最终是要放到M上执行的。
        - P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停，运行后续的goroutine等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了就会去其他P的队列里抢任务。
        - 比用操作系统调度更轻量
    - GO 语言中的操作系统线程和goroutine的关系：
        - 1.一个操作系统线程对应用户态多个goroutine
        - 2.go程序可以同时使用多个操作系统线程
        - 3.goroutine 和OS线程是多对多的关系 即 m:n
- m:n 把m个goroutine分配给n个操作系统线程去执行
- goroutine初始占用大小2K,比操作系统创建线程占用少
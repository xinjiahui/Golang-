## 并发和并行
- 并发：同一时间段内执行多个任务
- 并行：同一时刻执行多个任务

## Go语言的并发
- 通过goroutine 实现，goroutine类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个goroutine并发工作，goroutine是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。
- Go 语言还提供channel在多个goroutine间进行通信。goroutine和channel是Go语言秉承的CSP并发模式的重要实现基础。
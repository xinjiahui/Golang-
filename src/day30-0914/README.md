
## 总结
### time
- 时间类型
    - time.Time : time.Now()
    - 时间戳：
        - time.Now().Unix() // 1970.1.1 到现在的秒数
        - time.Now().UnixNano() // 1970.1.1 到现在的纳秒

- 时间间隔类型 
    - time.Duration :
        - time.Second()
        - time.Nanosecond()
- 时间操作
    - 时间对象 +/- 一个时间间隔对象
    ```
    //时间间隔
	fmt.Println(time.Second)
	// now +24小时
	fmt.Println(now.Add(24 * time.Hour))
	//Sub 两个时间相减
	nextYear, err := time.Parse("2006-01-02 15:04:05", "2022-09-04 20:52:00")
	if err != nil {
		fmt.Printf("parse time failed,err:%v", err)
	}
	//now = now.UTC()
	//nextYear = nextYear.UTC()
	d := nextYear.Sub(now)
	fmt.Println(d)
    ```
    - after/before

- 时间格式化
    ```
    fmt.Println(now.Format("2006/01/02 15:04:05"))
    ```
- 定时器
    ```
    timer := time.Tick(time.Second)
    for t := range timer{
        fmt.Println(t) //1秒钟执行一次
    }
    ```
 
- 解析字符串格式的时间（时区）
    ```
	//按照指定格式解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2022-09-04 21:35:55")
	//按照东八区的时区和格式去解析一个字符串格式的时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("loca location failed,err:%v", err)
		return
	}
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-09-04 21:35:55", loc)
	if err != nil {
		fmt.Printf("parse time failed,err:%v", err)
		return
	}
	fmt.Println(timeObj)
    ```
### 日志库
- time包
- 文件操作
- runtime.Caller()
- 结构体...

### 反射
- 接口类型的变量底层是分为两部分：动态类型和动态值
- 反射的应用：json等数据解析\ORM等工具...
- 反射的两个方法： 
    ```
    reflect.TypeOf  
    reflect.ValueOf
    ```
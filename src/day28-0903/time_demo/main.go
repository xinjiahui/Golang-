package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	//time.Unix()
	ret := time.Unix(1662203359, 0)
	fmt.Println(ret.Year())
	fmt.Println(ret.Date())
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
	fmt.Println("----------------")
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t) //	一秒钟执行一次
	//}
	//时间格式化:把语言中的时间对象，转化成字符串类型的时间
	//Y m d H M s
	//2006 1 2 15 04 	05
	// 2022-09-03
	fmt.Println(now.Format("2006-01-02"))
	// 2022/09/03 11:59:02
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2022/09/03 11:02:14 pm
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	// 2022/09/03 08:12:00.00 PM
	fmt.Println(now.Format("2006/01/02 03:04:05.000 PM"))
	//按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-08-03")
	if err != nil {
		fmt.Printf("parse time failed,err:%v", err)
	}
	fmt.Println(timeObj)
	//睡5s
	time.Sleep(5 * time.Second)
	fmt.Println("5s过去了...")
	fmt.Println(timeObj.Unix())
}
func f2() {
	now := time.Now() //本地时间
	fmt.Println(now)
	//明天的这个时间
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
	//时间相减
	td := timeObj.Sub(now)
	fmt.Println(td)

}

// 时间
func main() {
	f1()
	f2()
}

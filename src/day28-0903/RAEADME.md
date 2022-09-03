## 回顾
### 包
- 包的定义--->package关键字，包名通常是和目录名一致，不能包含-
    - 一个文件夹，就是一个包
    - 文件夹里面都是放的 .go文件
- 导入包--->import
    - 单行导入
    - 多行导入
    - 给导入的包起别名
    - 匿名导入---> sql包导入时使用
    - 包导入路径是从 $GOPATH/src 后面开始写
- 包中标识符（变量名/函数名/结构体/接口/常量...）可见性--->标识符首字母大写表示对外可见
- init（）
    - 包导入时候会自动执行
    - 一个包里面只有一个init（）
    - init（）没有参数也没有返回值也不能调用它
    - 多个包的init执行顺序
    - 一般用于做一些初始化操作...
### 接口
- 接口是一种抽象的类型
- 接口就是要实现方法的清单
- 接口的定义
```
type mover interface{
    run() //方法的签名：参数 返回值
    eat()
}
```
- 接口的实现
    - 实现了接口的所有方法就实现了这个变量
    - 实现了接口就可以当成这个接口的变量
- 接口变量
    - 实现了一个万能的变量可以保存所有实现了我这个接口的类型的值
    - 通常是作为函数的参数出现
- 空接口
    - 接口中没有定义任何方法，也就是说任意类型都实现了接口-->任何类型的变量都可以存到空接口中
    ```
    interface{}
    ```
    - 作为函数参数--->fmt.Println()
    - map[string]interface{}   
- 接口底层
    - 动态类型
    - 动态值
- 类型断言
    - 做类型断言的前提是：一定要是一个接口类型的变量
    ```
	//1.x.(T)
	v, ok := a.(int8)
	if ok {
		fmt.Println("猜对了,a是int8", v)
	} else {
		fmt.Println("猜错了,不是int8")
	}

	//2.switch
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case string:
		fmt.Println("string", v2)
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("滚～")
	}
    ```
### 文件操作

- 打开文件和关闭文件
    ```
    func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()

}
    ```
- 读文件(三种方法) （read bufio iouil）封装的越来越高级，越来越不灵活
    ```
    func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("Open file faild,err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	//读文件
	var tmp [128]byte //指定读的长度
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file faild,err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

//利用bufio这个包来读取文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed.err:%v", err)
		return
	}
	//记得关闭对象
	defer fileObj.Close()
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v", err)
			return
		}
		fmt.Print(line)
	}

}
func readFromFileByIouil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed,err:%v", err)
		return
	}
	fmt.Println(string(ret))

}
    ```
- 写文件(三种)

```
func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) //逻辑或操作，有一个为1，就为1
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	//write
	fileObj.Write([]byte("you dian yi huo\n"))
	//writestring
	fileObj.WriteString("哈哈哈哈")

	fileObj.Close()
}
func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) //逻辑或操作，有一个为1，就为1
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("你好呀～\n") //写到缓存中
	wr.Flush()               //将缓存中的内容写入文件
}
func writeDemo3() {
	str := "等一朵花开～"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file failed,err:%v", err)
		return
	}

}
```
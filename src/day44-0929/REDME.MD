### 单元测试
#### 单元测试
xxx/ccc.g0


#### 基准测试/性能测试
- 函数格式
```
func BenchmarkSplit(b  *testing.B){
    //b.N ：被测试函数执行的次数
}
```


### pprof

### flag标准库
- 声明变量
```
./xxx -name="辛佳慧" -age=18
```
    - 方法1:
``` 
nameVal1:=flag.String("name","xinjiahui","输入你的名字～") //返回的是一个指针变量
```
    - 方法2:
```
var nameVal2 string
flag.StringVar(&nameVal2,"name","xinjiahui","请输入你的名字")
```

- 必须要调用：
```
flag.Parse() //解析命令传入的参数，赋值给对应的变量
```
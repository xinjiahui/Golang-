[toc]
### 结构体
- 基本数据类型:表示现实中的物件有局限性
```
var name = "辛佳慧"
```
- 编程是代码解决现实生活中的问题，结构体也是一种数据类型，一种我们自己造的可以保存多个维度的类型
```
type person struct{
    name string
    age int
}

```
#### 匿名结构体
- 多用于临时场景
var a = struct{
    x int
    y int
}{10,20}
### 构造函数
    ``` 
    type person struct {
    	name string
    	age  int
    }
    //构造（构造结构体变量的）函数,返回值就是对应的结构体类型
    func newPerson(n string, i int) person{
    	return person{
    		name: n,
    		age:  i,
    	}
    }
    ```
### 方法和接收者
- 方法是有接收者的函数，接受者指的是哪个类型的变量可以调用这个函数
    ```
    type person struct {
    	name string
    	age  int
    }
    //方法
    func (p person) dream(str string) {
	    fmt.Printf("%s的梦想是学好Go语言%s", p.name, str)
    }
    func main(){
        p1.person{"zhangjie",18}
        p1.dream("test") 
    }
    ```

    - 结构体的嵌套
    - 结构体的匿名字段
    - JSON序列化与反序列化
    - 

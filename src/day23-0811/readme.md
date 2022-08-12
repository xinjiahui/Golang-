[toc]
### 结构体
- 结构体是值类型（区别于引用类型）
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
### 结构体的嵌套
    ```
    type addr struct{
	    provice string
	    city string
    }
    type student struct{
	    name string
	    address addr //嵌套别的结构体  名称  类型
    }
    ```
### 结构体的匿名字段
    ```
    type addr struct{
	    provice string
	    city string
    }
    type student struct{
	    name string
	    addr //匿名嵌套别的结构体，就使用类型名做名称
    }
    ```
### JSON序列化与反序列化
- 经常出现的问题：
    - 结构体内部的字段要大写！！！不大写别人时当问不到的
#### 序列化
    ```
	type point struct {
		X int `json:"x"` //
		Y int `json:"y"`
	}
	p2 := point{100, 200}
	b, err := json.Marshal(p2) //gong语言中，把错误当成值返回，错误通常作为第二个参数
	if err != nil {
		fmt.Printf("Marshal failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
    ```
#### 反序列化
    ```
	type point struct {
		X int `json:"x"` //
		Y int `json:"y"`
	}
	str1 := `{"xinjiahui":10,"zhangjie":20}`
	var po2 point //造一个结构体变量，准备接收反序列化的值
    //反序列化：由字符串--->Go语言中的结构体变量
	//反序列化时要传递指针
	err = json.Unmarshal([]byte(str1), &po2)
	if err != nil {
		fmt.Printf("unmarshal faild,err:%v\n", err)
	}
	fmt.Println(po2)
    ```

### 自定义类型
- 类型别名只在代码编写过程中有效，编译完之后就不存在，内置的byte和rune都属于类型别名
```
tpye MyInt int //自定义类型
```
### 类型别名
```
type newInt = int //类型别名
```
### 单元测试和基准测试
- Go语言中的测试依赖go test 命令，编写测试代码和编写普通的Go代码过程是类似的，并不要学习新的语法，规则或工具。

### 测试组
```
go test -run=TestSplit/case_3
```
### 测试覆盖率

```
go test -cover -coverprofile=c.out

go tool cover -html=text.out
```
### 基准测试
```
func BenchmarkName(b *testing.B){
    //
}

go test -bench=Split -benchmem
```

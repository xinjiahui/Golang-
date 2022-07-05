package main

import "fmt"

//递归：如果一个函数在内部调用自己，这个函数就是递归函数
// 递归是恶化处理那种问题相同.问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件
// 计算n的阶乘
// n!=n * (n-1)!
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
func main() {
	ret := f(5)
	ret1 := fc(5)
	fmt.Println(ret)
	fmt.Println(ret1)
}

// 上台阶：有n阶台阶，一次可以走1步，也可以走2步。有多少种走法
//如果有1个台阶，只有1种走法
//如果有2个台阶，有2中走法

func fc(n uint64) uint64 {
	if n <= 3 {
		return n
	}

	return fc(n-1) + f(n-2)
}

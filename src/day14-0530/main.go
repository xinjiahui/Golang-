package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Eizabeth.
分配规则如下：
a.名字中包好一个‘e’或‘E’
b.每个名字中包含一个‘i’或‘I’分2个金币
c.每个名字中包含一个‘o’或‘O’分3枚金币
d.每个名字中包含一个‘u’或‘U’分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现dispatchCoin函数。
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Eizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispathCoin()
	fmt.Println("剩下： ", left)
}

func dispathCoin() {
	//1.依次拿到每个人的名字
	//2.拿到一个人名根据分金币的规则去分金币
	//2.1每个人分的金币数应该保存到distribution中
	//2.2还要记录下剩余的金币数
	//3.整个第二步执行完就能得到最终每个人的金币数和剩余金币数
	return
}

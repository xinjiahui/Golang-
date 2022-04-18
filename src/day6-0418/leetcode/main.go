package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.判断字符串中汉字的数量
	//难点是判断一个人字符串是汉字
	var count int
	s1 := "hello 上海"
	//1.依次拿到字符串中的字符
	//2.判断当前这个字符是不是汉字
	//3.把汉字出现的次数累加得到最终结果
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++

		}
	}
	fmt.Println(count)
	// 2.how do you do 单词出现的次数
	//2.1 把字符串按照空格切割得到切片
	//2.2 遍历切片存储到一个map
	//2.3 累加出现的次数
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		// 1.如果原来map中不存在w这个key那么出现次数=1
		//2.如果map中存在w这个key。那么出现次数 +1
		for key, value := range m1 {
			fmt.Println(key, value)
		}
	}
	//3.回文判断
	//字符串从左读和从右往左读是一样的，那么就是回文
	ss := "上海自来水来自海上"
	// 把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}

	}
	fmt.Println("是回文")

}

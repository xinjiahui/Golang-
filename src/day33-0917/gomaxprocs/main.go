package main

import (
	"fmt"
	"runtime"
	"sync"
)

//GOMAXPROCS
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(6)         //默认是CPU的逻辑核心数
	fmt.Println(runtime.NumCPU()) //看操作系统核心数
	wg.Add(2)
	a()
	b()
	wg.Wait()
}

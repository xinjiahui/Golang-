package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http client

func main() {
	resp, err := http.Get("http://127.0.0.1:9090/hello?name=xinjiahui&age=18")
	if err != nil {
		fmt.Printf("get url failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close() //一定要记得关闭连接
	//从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.body failed,err:%v\n", err)
		return
	}
	fmt.Print(string(b))

}

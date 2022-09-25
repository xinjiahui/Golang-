package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp client
func main() {
	// 1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dail 127.0.0.1:20000 failed,err:", err)
	}
	//var msg string
	reader := bufio.NewReader(os.Stderr)
	for {
		fmt.Print("请说话:")
		msg, _ := reader.ReadString('\n') //读到换行
		msg = strings.TrimSpace(msg)
		//fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

	conn.Close()
}

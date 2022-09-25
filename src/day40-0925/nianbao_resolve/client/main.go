package main

import (
	proto "Golang-/src/day40-0925/nianbao_resolve/protocol"
	"fmt"
	"net"
)

// 粘包 client
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed,err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello,Hello,xinjiahui!`
		//调用协议编码数据
		b, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode failed err:", err)
		}
		conn.Write([]byte(b))
	}
}

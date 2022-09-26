package main

import (
	"fmt"
	"net"
	"strings"
)

//udp server
func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen udp failed,err:", err)
		return
	}
	defer conn.Close()
	//不需要建立连接直接首发数据
	var data [1024]byte //数组
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from UDP failed,err:", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))

		//发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}
}

// One-to-one ChatRoom Client
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("conn fail...")
	}
	defer conn.Close()
	fmt.Println("connect sever successed")

	var msg string  //声明msg为字符串变量

	for {
		msg = ""  //初始化msg为空值
		fmt.Printf("say to sever: ")
		fmt.Scan(&msg)  //输入客户端向服务器端要发送的消息
		//fmt.Println(msg)
		//msg_write := []byte(msg)
		//conn.Write(msg_write)
		conn.Write([]byte(msg))  //信息转化成字节流形式并向服务器端发送
		//此处造成客户端程序端口堵塞
		//fmt.Println([]byte(msg))

		//等待服务器端发送信息回来
		data := make([]byte, 255)
		msg_read, err := conn.Read(data)
		if msg_read == 0 || err != nil {
			fmt.Println("err")
		}
		msg_read_str := string(data[0:msg_read])
		if msg_read_str == "close" {
			conn.Write([]byte("close"))
			break
		}

		fmt.Println("sever say:", msg_read_str)
	}
	conn.Close()
}

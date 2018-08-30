// One-to-one ChatRoom Sever
package main

import (
	"fmt"
	"net"
)

//var ConnMap map[string]*net.TCPConn

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8000")  //打开监听接口
	if err != nil { //如果有错误
		fmt.Println("sever error")
	}

	defer listen_socket.Close()  //延迟服务器端关闭
	fmt.Println("sever is wating ....")

	for {
		conn, err := listen_socket.Accept()  //监听客户端的端口
		if err != nil {
			fmt.Println("conn fail ...")
		}
		fmt.Println("connect client successed")  //显示服务器端连接成功

		var msg string  //声明msg为字符串变量

		for {
			//开始接收客户端发过来的消息
			msg = ""  //字符串msg初始为空
			data := make([]byte, 255)  //创建并声明数据变量，为255位
			msg_read, err := conn.Read(data)  //接收由客户端发来的消息，字节赋值给msg_read，err为错误
			if msg_read == 0 || err != nil {  //如果读取的消息为0字节或者有错误
				fmt.Println("err")
			}

			msg_read_str := string(data[0:msg_read])  //将msg_read_str的字节格式转化成字符串形式
			if msg_read_str == "close" {  //如果接收到的客户端消息为close
				conn.Write([]byte("close"))
				break
			}
			//fmt.Println(string(data[0:msg_read]))
			fmt.Println("client say: ", msg_read_str)  //接收客户端发来的信息

			fmt.Printf("say to client: ")  //提示向客户端要说的话
			fmt.Scan(&msg)  //输入服务器端要对客户端说的话
			//conn.Write([]byte("hello client\n"))
			//msg_write := []byte(msg)
			conn.Write([]byte(msg))  //把消息发送给客户端
			//此处造成服务器端的端口堵塞
		}
		fmt.Println("client Close\n")
		conn.Close()  //关闭连接
	}

}


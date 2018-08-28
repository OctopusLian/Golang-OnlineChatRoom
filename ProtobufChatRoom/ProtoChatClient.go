package main

import (
	"net"
	"fmt"
	"Golang-OnlineChatRoom/ProtobufChatRoom/protocol"
	"Golang-OnlineChatRoom/ProtobufChatRoom/github.com/golang/protobuf/proto"
	"log"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")  //打开监听端口
	if err != nil {
		fmt.Println("conn fail...")
	}
	defer conn.Close()
	fmt.Println("client connect server successed \n")

	var nickname string
	fmt.Println("请输入您在聊天室中要显示的昵称：")
	fmt.Scan(&nickname)
	fmt.Println("hello,welcome to online chat room :",nickname)
	conn.Write([]byte("nickname|" + nickname))

	go CliHandle(conn)

	for{
			var datamsg string
			fmt.Println("请输入你要群发的消息：")
			fmt.Scan(&datamsg)

			testmsg := &protocol.Conn_ToS{
			Nickname:proto.String(nickname),
			Msg:proto.String(datamsg),
		}
		data,err := proto.Marshal(testmsg)
		if err != nil{
			log.Fatal("marshaling error:",err)
		}

		conn.Write(data)

		if datamsg == "quit"{
			conn.Write([]byte(datamsg + nickname))
			break
		}
	}
}

//客户端接收消息
func CliHandle(conn net.Conn){
	for{
			msgdata := make([]byte,255)
			msgdata_read,err := conn.Read(msgdata)
			if msgdata_read == 0 || err != nil{
				break
			}
			fmt.Println(string(msgdata[0:msgdata_read]))
	}
}

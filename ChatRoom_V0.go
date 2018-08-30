package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

//主、广播、连接处理、客户写入

type clients chan<- string  //设置广播器，客户对外发送信息的通道
// chan<- T 只能向此类通道发送元素值，而不能从它那里接收元素值

var(
	entering = make(chan string,10)  //建立客户进入房间的通道
	leaving = make(chan string,10)  //建立客户离开房间的通道
	messages = make(chan string,100)  //建立发送客户消息的通道
)
func broadcaster{
	clients := make(map[clients]bool)  //将map中的客户信息值返回给 clients
	for {
		select {  //选择发送消息的通道
		case information := <-entering  //接收客户进入房间的信号
			clients[clients] = true

		case information := <-leaving  //接收客户离开房间的信号
			delete(clients,information)  //删除用户的名字和发送的消息
			close(information)  //关闭

		case clients := <-messages  //接收客户发送消息的信号
			for information := range clients{  //把消息中的所有内容广播给所有客户
			//迭代出一个数组或切片值中的每个元素、一个字符串中的每个字符
				imformation <- msg
			}

		}
	}
}

func handleConn(conn net.Conn){
	ch := make(chan string)  //创建对外发送客户消息的通道，返回的字符串存入ch中
	go clientWriter()  //创建客户写入的goroutine

	human := conn.RemoteAddr().String()  //写入客户的名字
	ch <- "Your name is " + human  //“你的名字是...”写入ch
	messages <- human + "has arrived room"  //加入房间前显示 “...已经加入房间
	entering <- ch //打开进入房间的通道

	input := bufio.NewScanner(conn)  //读取客户输入的内容，并更新
	for input.Scan(){
		messages <- human + ": " + input.Text()  //在说的每句话之前加上客户的名字
	}

	leaving <- ch  //打开离开房间的通道
	messages <- who + "has left room"  //离开房间后显示 “...已经离开房间”
	conn.Close()  //客户离开房间后通道关闭


}

func clientWriter(conn net.Conn,ch <- chan string){
	for msg := range ch{
		fmt.Fprintln(conn,msg)  //将客户写入的所有内容记录在msg中并输出
	}
}

func main() {
	listener,err := net.Listen("tcp","127.0.0.1:8085")  //构建TCP协议 服务端程序的第一步
	if err != nil{  //判断error类型的值是否为nil
		log.Fatal(err)
		return
	}

	go broadcaster()  //并发执行广播的goroutine
	for{
		conn,err := listener.Accept()  //等待客户端的连接请求
		if err != nil{
			log.Print(err)
			continue
		}
		go handleConn(conn)  //对客户的连接进行处理的goroutine
	}


}

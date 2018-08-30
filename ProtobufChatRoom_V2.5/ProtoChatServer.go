package main

import (
	"net"
	"fmt"
)

var ConnMap map[int]net.Conn = make(map[int]net.Conn)

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8000")  //打开监听接口
	if err != nil {
		fmt.Println("server start error")
	}

	defer listen_socket.Close()
	fmt.Println("server is wating ....")
	var key int = 1
	for{
		conn,err := listen_socket.Accept()
		if err != nil{
			fmt.Println("connect failed ...")
		}
		fmt.Println(conn.RemoteAddr(),"connect successed !")
		ConnMap[key] = conn
		key++

		go ServHandle(conn)
	}

}

func ServHandle(conn net.Conn){


	for{
		datamsg := make([]byte,255)
		datamsg_read ,err := conn.Read(datamsg)

		if datamsg_read == 0 || err != nil{
			continue
		}
		fmt.Println(datamsg[0:datamsg_read])
		for _,v := range ConnMap{
			v.Write(datamsg[0:datamsg_read])
		}
	}
}

package main

import (
	"net"
	"fmt"
)

var ConnMap map[string]net.Conn = make(map[string]net.Conn)

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8000")  //打开监听接口
	if err != nil {
		fmt.Println("server start error")
	}

	defer listen_socket.Close()
	fmt.Println("server is wating ....")

	for{
		conn,err := listen_socket.Accept()
		if err != nil{
			fmt.Println("connect failed ...")
		}
		fmt.Println(conn.RemoteAddr(),"connect successed !")

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
		fmt.Println(datamsg)
		/*
		newtestmsg := &protocol.Conn_ToS{}
		err = proto.Unmarshal(datamsg,newtestmsg)
		if err != nil{
			log.Fatal("unmarshling error:",err)
		}*/
		/*
		data := &protocol.Conn_ToC{
			Nickname:proto.String(nickname),
			Msg:proto.String(),
		}
		data,err := proto.Marshal(testmsg)
		if err != nil{
			log.Fatal("marshaling error:",err)
		}*/
		conn.Write(datamsg)
	}
}

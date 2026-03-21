package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("hello")
	fmt.Println("Listening on port :6379")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		resp := newResp(conn)

		value, err := resp.Read()
		if err != nil{
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		conn.Write([]byte("+Ok\r\n"))
	}
}
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener,error := net.Listen("tcp","localhost:8000")
	if error != nil {
		log.Fatal(error)
	}

	for  {
		fmt.Println("Accept")
		conn,error := listener.Accept()
		if error != nil {
			log.Fatal(error)
		}

		go hanle(conn)

	}

}

func hanle(conn net.Conn)  {
	defer closeConn(conn)
	for  {
		_,error := io.WriteString(conn,"chenzhen")
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println("1111111111111111")
		time.Sleep(time.Second)
	}

}

func closeConn(conn net.Conn)  {
	fmt.Println("closeConn")
	conn.Close()

}
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn,error := net.Dial("tcp","localhost:8000")
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(conn)
	log.Println("connected")
	io.Copy(os.Stdout,conn)
}
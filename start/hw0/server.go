package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	connName := conn.RemoteAddr().String()

	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		fmt.Printf("can not answer to client %s\n", connName)
		return
	}
	fmt.Printf("answered to client %s\n", connName)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept error: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

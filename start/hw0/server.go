package main

import (
	"fmt"
	"net"
	"time"
)

const (
	network          = "tcp"
	address          = "localhost:8080"
	timeout          = time.Second * 5
	expectedResponse = "OK\n"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	connName := conn.RemoteAddr().String()

	_, err := conn.Write([]byte(expectedResponse))
	if err != nil {
		fmt.Printf("Ошибка отправки ответа клиенту %s\n", connName)
		return
	}
	fmt.Printf("Ответ клиенту %s\n", connName)
}

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

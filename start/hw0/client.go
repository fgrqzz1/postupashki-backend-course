package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

const (
	network          = "tcp"
	address          = "localhost:8080"
	timeout          = time.Second * 5
	expectedResponse = "OK\n"
)

func main() {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		fmt.Printf("Ошибка в подключении: %v\n", err)
		return
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Ошибка чтения: %v\n", err)
			return
		}

		if string(buf[:n]) != expectedResponse {
			fmt.Printf("Ошибка: получен неизвестный ответ: %v\n", string(buf[:n]))
			//panic("unexpected response")
		}
		fmt.Println("OK")
	}

}

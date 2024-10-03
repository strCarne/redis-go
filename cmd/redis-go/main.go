package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
		}

		buf := make([]byte, 216)
		bytesRead, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading: ", err.Error())
			continue
		}

		received := string(buf[:bytesRead])
		log.Println(received)

		conn.Write([]byte("+PONG\r\n"))
	}
}

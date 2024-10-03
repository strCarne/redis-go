package main

import (
	"bytes"
	"log"
	"net"
)

const (
	DefaultBufferSize = 256
)

func main() {
	//nolint:gosec
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Println("Failed to bind to port 6379")

		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
		}

		buf := make([]byte, DefaultBufferSize)

		bytesRead, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading: ", err.Error())

			continue
		}

		buf = buf[:bytesRead]

		received := string(buf)
		log.Println(received)

		for range bytes.Split(buf, []byte("\n")) {
			_, err := conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				log.Println("Error writing: ", err.Error())
			}
		}
	}
}

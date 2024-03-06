package main

import (
	"fmt"
	"net"
)

func main() {
	tcp, err := net.Listen("tcp", ":25565")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := tcp.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
}

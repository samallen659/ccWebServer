package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	svr, _ := net.ListenTCP("tcp", addr)
	for {
		conn, err := svr.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()

	b := make([]byte, 4096)
	_, err := conn.Read(b)
	if err != nil {
		//TODO: write http error response
		return
	}

	reqStr := string(b)
	fLine := strings.Split(reqStr, "\n")[0]
	fmt.Println(fLine)
}

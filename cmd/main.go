package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	svr, _ := net.ListenTCP("tcp", addr)
	for {
		conn, err := svr.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		b, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read conn: %s", err.Error())
		}
		fmt.Println(string(b))
		conn.Close()
	}
}

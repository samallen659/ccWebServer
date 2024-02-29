package server

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	addr *net.TCPAddr
}

func NewServer(addrStr string) (*Server, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		return nil, err
	}

	return &Server{addr}, nil
}

func (s *Server) Listen() error {
	svr, err := net.ListenTCP("tcp", s.addr)
	if err != nil {
		return err
	}
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

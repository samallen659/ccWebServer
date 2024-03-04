package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	addr   *net.TCPAddr
	router *Router
}

func NewServer(addrStr string, wwwPath string) (*Server, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		return nil, err
	}

	r, err := NewRouter(wwwPath)
	if err != nil {
		return nil, err
	}

	return &Server{addr, r}, nil
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
		log.Println(err)
		//TODO: write http error response
		return
	}

	reqStr := string(b)

	req, err := NewRequest(reqStr)
	if err != nil {
		log.Println(err)
		//TODO: write http error response
		return
	}

	res := NewResponse()
	res.SetStatus(HTTP_OK)
	res.SetBody(fmt.Sprintf("Requested Path: %s\n", req.Header.Path))

	conn.Write(res.Marshall())
}

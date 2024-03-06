package server

import (
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

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn *net.TCPConn) {
	defer conn.Close()
	res := NewResponse()

	b := make([]byte, 4096)
	_, err := conn.Read(b)
	if err != nil {
		log.Println(err)
		res.SetStatus(HTTP_INTERNAL_SEVER_ERROR)
		conn.Write(res.Marshall())
		return
	}

	reqStr := string(b)

	req, err := NewRequest(reqStr)
	if err != nil {
		log.Println(err)
		res.SetStatus(HTTP_BAD_REQUEST)
		conn.Write(res.Marshall())
		return
	}

	s.router.RouteRequest(req, res)

	conn.Write(res.Marshall())
}

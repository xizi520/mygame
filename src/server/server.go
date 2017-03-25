package server

import (
	"net"
	"strconv"
)

type TcpServer struct {
	listener *net.TCPListener
	Name string
	Sessions map[int]interface{}
}

func (server *TcpServer) Serve(port int) error {
	addr, err := net.ResolveTCPAddr("tcp", strconv.Itoa(port))
	if err != nil {
		return err
	}

	server.listener net.ListenTCP("tcp", addr)
}

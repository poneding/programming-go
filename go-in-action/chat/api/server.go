package api

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	IP   string
	Port int

	OnlineUserMap map[string]*User
	mapLock       sync.RWMutex
	Message       chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		IP:   ip,
		Port: port,

		OnlineUserMap: make(map[string]*User),
		Message:       make(chan string),
	}
}

func (s *Server) Start() {
	listener, e := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if e != nil {
		panic(e)
	}

	defer listener.Close()

	go s.ListenMessager()

	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("listener accept err:", e)
			continue
		}

		go s.Handler(conn)
	}

}

func (s *Server) Handler(conn net.Conn) {
	// fmt.Println("OK")

	s.mapLock.Lock()
	u := NewUser(conn)

	s.OnlineUserMap[u.Name] = u
	s.mapLock.Unlock()

	s.BroadCast(u, "online")

	select {}
}

func (s *Server) BroadCast(u *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]-%s: %s", u.Addr, u.Name, msg)
	s.Message <- sendMsg
}

func (s *Server) ListenMessager() {
	for {
		msg := <-s.Message
		s.mapLock.Lock()
		for _, u := range s.OnlineUserMap {
			u.C <- msg
		}
		s.mapLock.Unlock()
	}
}

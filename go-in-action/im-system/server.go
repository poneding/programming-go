package goim

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

// Server The server
type Server struct {
	IP   string
	Port int

	sync.RWMutex
	OnlineUsersMap map[string]*User
	MessageChan    chan Message
}

// NewServer Create a new server
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:             ip,
		Port:           port,
		OnlineUsersMap: make(map[string]*User),
		MessageChan:    make(chan Message),
	}
}

// HandleConn Handle new connection
func (s *Server) HandleConn(conn net.Conn) {
	u := NewUser(conn, s)
	go u.ListenMsg()

	u.Online()

	// Accept client messages
	go func() {
		buf := make([]byte, 4096) // what's if the message is larger than 4096?
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				u.Offline()
				return
			}

			if err != nil && err != io.EOF {
				log.Println("Read from conn err:", err)
				return
			}

			msg := string(buf[:n-1]) // remove '\n' ??
			u.DoMsg(msg)

			// keep user alive status
			u.Alive <- struct{}{}
		}
	}()

	// Block? why block here?
	for {
		select {
		case <-u.Alive:
			// Do nothing just donates that the user is alive
		case <-time.After(time.Minute * 10):
			// Force to kick out the user if no activity for 10 seconds
			s.safeDeleteOnlineUser(u)
			close(u.Alive)
			conn.Close()

			return
		}
	}
}

// BroadcastMsg Send a message to broadcast channel
func (s *Server) BroadcastMsg(u *User, msg string) {
	log.Println(fmt.Sprintf("(%s): %s", u.Name, msg))
	s.MessageChan <- Message{
		Sender:  u.Name,
		Content: msg,
	}
}

func (s *Server) safeAddOnlineUser(u *User) {
	s.Lock()
	s.OnlineUsersMap[u.Name] = u
	s.Unlock()
}

func (s *Server) safeDeleteOnlineUser(u *User) {
	s.Lock()
	delete(s.OnlineUsersMap, u.Name)
	s.Unlock()
}

// ListenMsg Listen for new messages and broadcast them to all users
func (s *Server) ListenMsg() {
	for {
		msg := <-s.MessageChan

		s.RLock()
		for _, u := range s.OnlineUsersMap {
			u.IncomeMsgChan <- msg
		}
		s.RUnlock()
	}
}

// Start the server
func (s *Server) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	// Start listening for new messages
	go s.ListenMsg()

	for {
		// Accept new connection
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("Error accepting new connection: %v", err)
			continue
		}

		// Handle new connection
		go s.HandleConn(conn)
	}
}

package goim

import (
	"fmt"
	"github.com/google/uuid"
	"net"
	"regexp"
	"strings"
)

var (
	sysUser = &User{
		Name: "system",
		Addr: "SYSTEM",
	}
)

// User The user
type User struct {
	Name string
	Addr string

	IncomeMsgChan chan Message
	conn          net.Conn
	server        *Server
	Alive         chan struct{}
}

// NewUser Create a new user
func NewUser(conn net.Conn, s *Server) *User {
	addr := conn.RemoteAddr().String()
	u := &User{
		Name: uuid.NewString()[:8], // uuid name
		Addr: addr,

		IncomeMsgChan: make(chan Message),
		conn:          conn,
		server:        s,
		Alive:         make(chan struct{}),
	}
	return u
}

// Online Online
func (u *User) Online() {
	u.server.safeAddOnlineUser(u)

	u.server.BroadcastMsg(u, "I'm in")
}

// Offline Offline
func (u *User) Offline() {
	u.server.safeDeleteOnlineUser(u)

	u.server.BroadcastMsg(u, "I'm out")
}

// chat format: @username:msg
var chatreg = regexp.MustCompile(`^@`)

// DoMsg Do message
func (u *User) DoMsg(msg string) {
	if action, ok := msgActionsMap[msg]; ok {
		action(u, Message{
			Sender:  u.Name,
			Content: msg,
		})
		return
	}

	if strings.HasPrefix(msg, "rename ") {
		newName := strings.Split(msg, "rename ")[1]
		if newName == u.Name {
			return
		}

		if _, ok := u.server.OnlineUsersMap[newName]; ok {
			u.conn.Write([]byte("The name has been taken\n"))
			return
		}

		u.server.Lock()
		delete(u.server.OnlineUsersMap, u.Name)
		u.Name = newName
		u.server.OnlineUsersMap[u.Name] = u
		u.server.Unlock()

		return
	}

	if strings.HasPrefix(msg, "@") && strings.ContainsRune(msg, ' ') {
		msgParts := strings.SplitN(msg[1:], " ", 2)
		to := msgParts[0]
		content := msgParts[1]

		if to == u.Name {
			u.AcceptMsg(sysUser, "You can't send message to yourself")
		}
		if content == "" {
			u.AcceptMsg(sysUser, "Message can't be empty")
		}

		if receiver, ok := u.server.OnlineUsersMap[to]; !ok {
			u.AcceptMsg(sysUser, fmt.Sprintf("User %s is not online", to))
		} else {
			receiver.AcceptMsg(u, content)
		}
		return
	}

	u.server.BroadcastMsg(u, msg)
}

func (u *User) AcceptMsg(sender *User, msg string) {
	u.IncomeMsgChan <- Message{
		Sender:  sender.Name,
		Content: msg,
	}
}

// ListenMsg Listen for new messages
func (u *User) ListenMsg() {
	for {
		msg := <-u.IncomeMsgChan
		if msg.Sender != u.Name {
			u.conn.Write([]byte(fmt.Sprintf("[%s]: %s\n", msg.Sender, msg.Content)))
		}
	}
}

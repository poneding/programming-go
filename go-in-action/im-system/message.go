package goim

import "fmt"

type Message struct {
	Sender  string
	Content string
}

type msgAction func(u *User, msg Message)

var (
	who = func(u *User, msg Message) {
		u.server.RLock()
		onlineUsers := make([]string, 0, len(u.server.OnlineUsersMap))
		for _, u := range u.server.OnlineUsersMap {
			onlineUsers = append(onlineUsers, u.Name)
		}
		u.server.RUnlock()
		//s.BroadcastMsg(sysUser, fmt.Sprintf("Online users(%d): %v", len(onlineUsers), onlineUsers))
		u.AcceptMsg(sysUser, fmt.Sprintf("Online users(%d): %v", len(onlineUsers), onlineUsers))
	}

	msgActionsMap = map[string]msgAction{
		"who": who,
	}
)

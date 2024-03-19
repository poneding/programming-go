package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

// 不可用，客户端连接返回 403
// 使用 gorilla/websocket 库
func main() {
	e := gin.Default()
	e.GET("/ws", func(c *gin.Context) {
		websocket.Handler(func(conn *websocket.Conn) {
			defer conn.Close()
			for {
				// 写回客户端消息
				err := websocket.Message.Send(conn, "hello client")
				if err != nil {
					fmt.Println("send message error:", err)
					return
				}

				// 读取客户端消息
				var msg string
				err = websocket.Message.Receive(conn, &msg)
				if err != nil {
					fmt.Println("receive message error:", err)
					return
				}
				fmt.Println("receive message:", msg)
			}
		}).ServeHTTP(c.Writer, c.Request)
	})

	e.Run(":8080")
}

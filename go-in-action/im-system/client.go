package goim

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type client struct {
	serverIP   string
	serverPort int
	conn       net.Conn
}

func NewClient(serverIP string, serverPort int) *client {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		log.Fatalln("NewClient err:", err)
	}
	return &client{
		serverIP:   serverIP,
		serverPort: serverPort,
		conn:       conn,
	}
}

func (c *client) Run() {
	go c.AcceptResponse()

	for {
		var action int8
		fmt.Println("Please select the action you want to do:")
		fmt.Println("0. Exit")
		fmt.Println("1. Chat")
		fmt.Println("2. Rename")
		fmt.Println("3. Announce")
		fmt.Scanln(&action)

		switch action {
		case 0:
			c.Exit()
		case 1:
			c.Chat()
		case 2:
			c.Rename()
		case 3:
			c.Announce()
		default:
			fmt.Println("Invalid action")
		}
	}
}

func (c *client) AcceptResponse() {
	io.Copy(os.Stdout, c.conn)
}

func (c *client) GetOnlineUsers() {
	msg := "who\n"
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		log.Fatalln("GetOnlineUsers err:", err)
	}
}

// Announce to all users
func (c *client) Announce() {
	//var msg string
	fmt.Println("Please input the message you want to announce('exit' to exit):")
	msg := readInput()

	for msg != "exit" {
		if msg != "" {
			_, err := c.conn.Write([]byte(msg + "\n"))
			if err != nil {
				log.Fatalln("Announce err:", err)
			}
		}

		fmt.Println("Please input the message you want to announce('exit' to exit):")
		msg = readInput()
	}
}

var stdinReader *bufio.Reader

func readInput() string {
	if stdinReader == nil {
		stdinReader = bufio.NewReader(os.Stdin)
	}
	line, _, _ := stdinReader.ReadLine()
	return string(line)
}

// Chat Start chat with a user
func (c *client) Chat() {
	c.GetOnlineUsers()
	fmt.Println("Please input the user you want to chat with('exit' to exit):")
	to := readInput()

	for to != "exit" {
		fmt.Println("Please input the message you want to send('exit' to exit):")
		msg := readInput()

		for msg != "exit" {
			if msg != "" {
				msg = "@" + to + " " + msg + "\n"
				_, err := c.conn.Write([]byte(msg))
				if err != nil {
					log.Fatalln("Chat err:", err)
				}
			}

			fmt.Println("Please input the message you want to send('exit' to exit):")
			msg = readInput()
		}

		c.GetOnlineUsers()
		fmt.Println("Please input the user you want to chat with('exit' to exit):")
		to = readInput()
	}
}

// Rename client name
func (c *client) Rename() {
	fmt.Println("Please input the name you want to use:")
	name := readInput()
	msg := "rename " + name + "\n"
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		log.Fatalln("Rename err:", err)
	}
	fmt.Println("Rename successfully, your new name is:", name)
}

// Exit client
func (c *client) Exit() {
	fmt.Println("Bye~")
	c.conn.Close()
	os.Exit(0)
}

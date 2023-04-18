package main

import (
	"chat/api"
)

func main() {
	s := api.NewServer("127.0.0.1", 8080)
	s.Start()
}

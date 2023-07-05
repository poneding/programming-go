package main

import "goim"

func main() {
	// Run the server
	goim.NewServer("127.0.0.1", 8081).Start()
}

package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, err := ParseToTCPAddr("1.1.1.1:1234")
	if err != nil {
		panic(err)
	}
	fmt.Println(tcpAddr.String())
}

func ParseToTCPAddr(v string) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp", v)
}

func ParseToUDPAddr(v string) (*net.UDPAddr, error) {
	return net.ResolveUDPAddr("udp", v)
}

func ParseToIPAddr(v string) (*net.IPAddr, error) {
	return net.ResolveIPAddr("ip", v)
}

func ParseToUnixAddr(v string) (*net.UnixAddr, error) {
	return net.ResolveUnixAddr("unix", v)
}

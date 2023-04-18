package main

import (
	"bytes"
	"log"
	"strings"
)

func main() {
	ioReader2Bytes()

	bytes2IOReader()
}

func ioReader2Bytes() {
	reader := strings.NewReader("hello world")

	buf := bytes.Buffer{}
	buf.ReadFrom(reader)

	bytesData := buf.Bytes()
	log.Println("bytesData:", string(bytesData))
}

func bytes2IOReader() {
	bytesData := []byte("hello world")

	reader := bytes.NewReader(bytesData)
	buf := make([]byte, 5)
	n, err := reader.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(buf[:n]))
}

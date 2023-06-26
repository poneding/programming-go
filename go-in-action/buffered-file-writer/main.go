package main

import (
	"log"
	"os"
	"time"
)

func main() {
	t := time.Now()
	// WriteDirectly() // 91ms
	WriteWithBuffer() // 1ms
	log.Println("elapsed:", time.Since(t).String())
}

func WriteDirectly() {
	fout, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	for i := 0; i < 10000; i++ {
		fout.Write([]byte("hello world\n"))
	}
}

func WriteWithBuffer() {
	fout, err := os.OpenFile("test2.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	bw := NewBufferedFileWriter(fout, 4096)
	defer bw.Flush()

	for i := 0; i < 10000; i++ {
		bw.Write([]byte("hello world\n"))
	}
}

type BufferedFileWriter struct {
	fout   *os.File
	buf    []byte
	offset int
}

func NewBufferedFileWriter(fout *os.File, bufSize int) *BufferedFileWriter {
	return &BufferedFileWriter{
		fout:   fout,
		buf:    make([]byte, bufSize),
		offset: 0,
	}
}

func (w *BufferedFileWriter) Write(b []byte) {
	if len(b) >= cap(w.buf) {
		w.Flush()
		w.fout.Write(b)
		return
	}

	if len(b) >= cap(w.buf)-w.offset {
		w.Flush()
	}
	copy(w.buf[w.offset:], b)
	w.offset += len(b)
}

func (w *BufferedFileWriter) Flush() {
	w.fout.Write(w.buf[0:w.offset])
	w.offset = 0
}

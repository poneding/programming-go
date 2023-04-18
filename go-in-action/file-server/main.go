package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	// go run main.go
	// 1、直接访问 localhost:8080，可以看到 files 目录下的文件列表
	http.Handle("/", http.FileServer(http.Dir("./files")))

	// 2、无法访问 localhost:8080/files，因为它会去找 ./files/files 目录
	// http.Handle("/files", http.FileServer(http.Dir("./files")))

	// 3、访问 localhost:8080/files，可以看到 files 目录下的文件列表
	// http.Handle("/files/", http.FileServer(http.Dir(".")))

	// 4、访问 localhost:8080/files files 目录下的文件列表
	// http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir("./files"))))
	http.ListenAndServe(":8080", nil)
}

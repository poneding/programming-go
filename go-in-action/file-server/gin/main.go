package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 使用标志库 net/http 实现文件服务
// func main() {
// 	// go run main.go
// 	// 1、直接访问 localhost:8080，可以看到 files 目录下的文件列表
// 	http.Handle("/", http.FileServer(http.Dir("./files")))

// 	// 2、无法访问 localhost:8080/files，因为它会去找 ./files/files 目录
// 	// http.Handle("/files", http.FileServer(http.Dir("./files")))

// 	// 3、访问 localhost:8080/files，可以看到 files 目录下的文件列表
// 	// http.Handle("/files/", http.FileServer(http.Dir(".")))

// 	// 4、访问 localhost:8080/files files 目录下的文件列表
// 	// http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir("./files"))))
// 	http.ListenAndServe(":8080", nil)
// }

type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	Email string                `form:"email" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/index", "./static")
	router.StaticFS("/files", http.Dir("./files"))
	router.POST("/upload", func(c *gin.Context) {
		var bindFile BindFile

		// Bind file
		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		// Save uploaded file
		file := bindFile.File
		dst := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, path.Join("./files", dst)); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		// c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully by %s(%s).", file.Filename, bindFile.Name, bindFile.Email))
		c.Data(http.StatusOK, "text/html", []byte(fmt.Sprintf("File <a href='/files/%s'>%s</a> uploaded successfully by %s(%s).", file.Filename, file.Filename, bindFile.Name, bindFile.Email)))
	})
	router.Run(":8080")
}

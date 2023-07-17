package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// Timer 接口计时中间件
func Timer() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("enter timer middleware")
		start := time.Now()
		c.Next()
		log.Printf("time consumed: %d ms", time.Since(start).Milliseconds())
		log.Println("leave timer middleware")
	}
}

var limitChan = make(chan struct{}, 10) // 限制并发数为 10

// Limiter 接口限流中间件
func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("enter limiter middleware")
		limitChan <- struct{}{}
		c.Next()
		<-limitChan
		log.Println("leave limiter middleware")
	}
}

func hello(c *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	c.String(http.StatusOK, "Hello World\n")
}

func main() {
	engine := gin.Default()

	engine.Use(Timer(), Limiter())
	engine.GET("/hello", hello)
	engine.GET("/double-hello", hello, hello)

	engine.Run(":8080")
}

package main

import (
	"fmt"
	"log"
	"os"
	"rabbitmq-sample"
)

var mqurl string

func main() {
	qname := "q0"
	mq := rabbitmq.NewWorkClient(mqurl, qname)
	defer mq.Close()
	for i := 0; i < 50; i++ {
		mq.Publish(fmt.Sprintf("Hello World! MSG_ID: %d", i))
	}
}

func init() {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPass := os.Getenv("RABBITMQ_PASS")
	if rabbitmqHost == "" {
		log.Fatalln("RABBITMQ_HOST environment variable is required")
	}
	if rabbitmqPort == "" {
		rabbitmqPort = "5672"
	}
	if rabbitmqUser == "" {
		log.Fatalln("RABBITMQ_USER environment variable is required")
	}
	if rabbitmqPass == "" {
		log.Fatalln("RABBITMQ_PASS environment variable is required")
	}
	mqurl = fmt.Sprintf("amqp://%s:%s@%s:%s/my_vhost", rabbitmqUser, rabbitmqPass, rabbitmqHost, rabbitmqPort)
}

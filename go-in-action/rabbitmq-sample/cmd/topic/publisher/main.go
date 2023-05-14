package main

import (
	"fmt"
	"log"
	"os"
	"rabbitmq-sample"
	"time"
)

var mqurl string

func main() {
	exchange := "e3"
	topicKey1 := "topic.k1"
	topicKey2 := "topic.k2"
	mq1 := rabbitmq.NewTopicClient(mqurl, exchange, topicKey1)
	mq2 := rabbitmq.NewTopicClient(mqurl, exchange, topicKey2)
	for i := 0; i < 1000; i++ {
		time.Sleep(100 * time.Millisecond)
		mq1.Publish(fmt.Sprintf("Hello World! MSG_ID: %d", i))
		if i%7 == 0 {
			mq2.Publish(fmt.Sprintf("Error message. MSG_ID: %d", i))
		}
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

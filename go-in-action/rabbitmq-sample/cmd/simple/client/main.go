package main

import (
	"fmt"
	"rabbitmq-sample"
)

func main() {
	mq := rabbitmq.NewSimpleRabbitMQClient("amqp://guest:guest@cloud.io:5672/", "test", "test")
	mq.Consume(func(msg string) {
		fmt.Println(msg)
	})
}

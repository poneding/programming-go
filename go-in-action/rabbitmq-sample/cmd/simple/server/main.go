package main

import (
	"rabbitmq-sample"

	"github.com/google/uuid"
)

func main() {
	mq := rabbitmq.NewSimpleRabbitMQClient("amqp://guest:guest@cloud.io:5672/", "test", "test")
	mq.Publish("Hello World! MSGID: " + uuid.NewString())
}

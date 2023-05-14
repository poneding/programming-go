package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RoutingClient struct {
	*client
}

func NewRoutingClient(url, exchange, routingKey string) *RoutingClient {
	return &RoutingClient{
		newClient(url, "", exchange, routingKey),
	}
}
func (mq *RoutingClient) Publish(msg string) {
	// 获取交换机
	err := mq.channel.ExchangeDeclare(
		mq.exchange,
		"direct", // direct, 单向
		true,     // durable，持久化
		false,    // auto-deleted，自动删除，当最后一个消费者断开连接之后，队列会自动删除
		false,    // internal，内部使用，如果为true，exchange不能被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,    // no-wait，不等待
		nil,
	)

	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mq.channel.PublishWithContext(
		ctx,
		mq.exchange,
		mq.key,
		false, // mandatory, 如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false, // immediate, 如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	failOnError(err, "Failed to publish a message")
}

func (mq *RoutingClient) Consume(consumeFunc func(msg string)) {
	// 获取交换机
	err := mq.channel.ExchangeDeclare(
		mq.exchange,
		"direct", // direct, 单向
		true,     // durable，持久化
		false,    // auto-deleted，自动删除，当最后一个消费者断开连接之后，队列会自动删除
		false,    // internal，内部使用，如果为true，exchange不能被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,    // no-wait，不等待
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	// 获取队列
	q, err := mq.channel.QueueDeclare(
		mq.qname, // qname为空
		true,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	// 绑定队列到交换机
	err = mq.channel.QueueBind(
		q.Name,
		mq.key,
		mq.exchange,
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := mq.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	log.Println("start to consume messages...")
	forever := make(chan struct{})
	go func() {
		for msg := range msgs {
			consumeFunc(string(msg.Body))
		}
	}()

	<-forever
}

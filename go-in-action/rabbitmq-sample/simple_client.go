package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 简单模式，一个生产者，一个消费者，消息只能被一个消费者消费
type SimpleClient struct {
	*client
}

func NewSimpleClient(url, qname string) *SimpleClient {
	// Simple 模式使用默认的exchange，routing key为队列名
	mq := newClient(url, qname, "", qname)

	return &SimpleClient{
		mq,
	}
}

func (mq *SimpleClient) Publish(msg string) {
	// 声明队列
	_, err := mq.channel.QueueDeclare(
		mq.qname, // name
		false,    // durable, 持久化
		false,    // auto-delete, 自动删除，当最后一个消费者断开连接之后，队列会自动删除
		false,    // exclusive, 排他性，如果为true，只能被创建者使用，且连接断开时自动删除
		false,    // no-wait，不等待，如果为true，不等待服务器响应，直接返回
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// 声明交换机
	// err = mq.rabbitmq.channel.ExchangeDeclare(
	// 	mq.rabbitmq.exchange, // name
	// 	"direct",             // type，direct、fanout、topic、headers，这里使用direct，direct类型的exchange根据routing key全文匹配去寻找队列
	// 	true,                 // durable, 持久化
	// 	false,                // auto-deleted, 自动删除，当最后一个消费者断开连接之后，队列会自动删除
	// 	false,                // internal, 内部使用，如果为true，exchange不能被client用来推送消息，仅用来进行exchange和exchange之间的绑定
	// 	false,                // no-wait, 不等待，如果为true，不等待服务器响应，直接返回
	// 	nil,                  // arguments
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to declare an exchange: %s", err)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mq.channel.PublishWithContext(
		ctx,
		mq.exchange, // exchange
		mq.key,      // routing key
		false,       // mandatory, 如果为true，当exchange发送消息到queue失败时，会将消息返回给发送者
		false,       // immediate, 如果为true，当exchange发送消息到queue失败时，不会将消息存入queue中，而是直接返回给发送者
		amqp.Publishing{
			// DeliveryMode: amqp.Persistent, // 持久化，如果为2，表示消息持久化，即使rabbitmq重启，消息也不会丢失
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish a message")
}

func (mq *SimpleClient) Consume(consumeFunc func(msg string)) {
	// Simple 模式使用默认的exchange，routing key为队列名

	// 声明队列
	_, err := mq.channel.QueueDeclare(
		mq.qname, // name
		false,    // durable, 持久化
		false,    // auto-delete, 自动删除，当最后一个消费者断开连接之后，队列会自动删除
		false,    // exclusive, 排他性，如果为true，只能被创建者使用，且连接断开时自动删除
		false,    // no-wait，不等待，如果为true，不等待服务器响应，直接返回
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := mq.channel.Consume(
		mq.qname, // queue
		"",       // consumer
		true,     // auto-ack, 自动确认，如果为true，当消费者收到消息后，会自动发送ack给rabbitmq，告诉rabbitmq这条消息已经被消费了，rabbitmq可以将消息从队列中删除
		false,    // exclusive, 排他性，如果为true，只能被创建者使用，且连接断开时自动删除
		false,    // no-local, 如果为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,    // no-wait，不等待，如果为true，不等待服务器响应，直接返回
		nil,      // arguments
	)
	failOnError(err, "Failed to register a consumer")

	log.Println("start to consume message...")
	forever := make(chan struct{})
	go func() {
		for msg := range msgs {
			consumeFunc(string(msg.Body))
		}
	}()
	<-forever
}

func (mq *SimpleClient) ConsumeWithMunal(consumeFunc func(msg string)) {
	// Simple 模式使用默认的exchange，routing key为队列名

	// 声明队列
	_, err := mq.channel.QueueDeclare(
		mq.qname, // name
		false,    // durable, 持久化
		false,    // auto-delete, 自动删除，当最后一个消费者断开连接之后，队列会自动删除
		false,    // exclusive, 排他性，如果为true，只能被创建者使用，且连接断开时自动删除
		false,    // no-wait，不等待，如果为true，不等待服务器响应，直接返回
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := mq.channel.Consume(
		mq.qname, // queue
		"",       // consumer
		false,    // auto-ack, 自动确认，如果为true，当消费者收到消息后，会自动发送ack给rabbitmq，告诉rabbitmq这条消息已经被消费了，rabbitmq可以将消息从队列中删除
		false,    // exclusive, 排他性，如果为true，只能被创建者使用，且连接断开时自动删除
		false,    // no-local, 如果为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,    // no-wait，不等待，如果为true，不等待服务器响应，直接返回
		nil,      // arguments
	)
	failOnError(err, "Failed to register a consumer")

	// 限流，一次只能消费一条消息
	err = mq.channel.Qos(
		1, // prefetchCount，当前消费者最大消费数
		0, // 服务器传递的最大容量
		false,
	)
	failOnError(err, "Failed to set Qos")

	log.Println("start to consume message...")
	forever := make(chan struct{})
	go func() {
		for msg := range msgs {
			consumeFunc(string(msg.Body))
			msg.Ack(false) // true 表示确认所有未确认的消息，false 表示只确认当前消息
		}
	}()
	<-forever
}

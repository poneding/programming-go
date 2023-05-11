package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// rabbitmq url, format: amqp://user:password@host:port/vhost
	mqurl    string
	mqname   string
	exchange string
	key      string
}

// type MSG struct {
// 	ID        any       `json:"id"`
// 	CreatedAt time.Time `json:"created_at"`
// 	Body      any       `json:"body"`
// }

type SimpleRabbitMQClient struct {
	rabbitmq *RabbitMQ
}

func newRabbitMQ(mqname, exchange, key string) *RabbitMQ {
	return &RabbitMQ{
		mqname:   mqname,
		exchange: exchange,
		key:      key,
	}
}

// connect 建立连接
func (mq *RabbitMQ) connect() {
	var err error
	mq.conn, err = amqp.Dial(mq.mqurl)
	failOnError(err, "Failed to connect to RabbitMQ")
	mq.channel, err = mq.conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func NewSimpleRabbitMQClient(url, exchange, key string) *SimpleRabbitMQClient {
	mq := newRabbitMQ("", exchange, key)

	return &SimpleRabbitMQClient{
		rabbitmq: mq,
	}

	// return &RabbitMQ{
	// 	mqurl:    url,
	// 	exchange: exchange,
	// 	key:      key,
	// }
}

func (mq *SimpleRabbitMQClient) Publish(msg string) {
	mq.rabbitmq.connect()
	defer mq.rabbitmq.Close()

	// Simple 模式使用默认的exchange，routing key为队列名

	// 声明队列
	_, err := mq.rabbitmq.channel.QueueDeclare(
		mq.rabbitmq.mqname, // name
		false,              // durable, 持久化
		false,              // auto-delete, 自动删除，当最后一个消费者断开连接之后，队列会自动删除
		false,              // exclusive, 排他性，如果为true，只能被创建者使用，且连接断开时自动删除
		false,              // no-wait，不等待，如果为true，不等待服务器响应，直接返回
		nil,                // arguments
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

	err = mq.rabbitmq.channel.Publish(
		mq.rabbitmq.exchange, // exchange
		mq.rabbitmq.key,      // routing key
		false,                // mandatory, 如果为true，当exchange发送消息到queue失败时，会将消息返回给发送者
		false,                // immediate, 如果为true，当exchange发送消息到queue失败时，不会将消息存入queue中，而是直接返回给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish a message")
}

func (mq *SimpleRabbitMQClient) Consume(consumeFunc func(msg string)) {
	mq.rabbitmq.connect()
	defer mq.rabbitmq.Close()

	// Simple 模式使用默认的exchange，routing key为队列名

	// 声明队列
	_, err := mq.rabbitmq.channel.QueueDeclare(
		mq.rabbitmq.mqname, // name
		false,              // durable, 持久化
		false,              // auto-delete, 自动删除，当最后一个消费者断开连接之后，队列会自动删除
		false,              // exclusive, 排他性，如果为true，只能被创建者使用，且连接断开时自动删除
		false,              // no-wait，不等待，如果为true，不等待服务器响应，直接返回
		nil,                // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	msgs, err := mq.rabbitmq.channel.Consume(
		mq.rabbitmq.mqname, // queue
		"",                 // consumer
		true,               // auto-ack, 自动确认，如果为true，当消费者收到消息后，会自动发送ack给rabbitmq，告诉rabbitmq这条消息已经被消费了，rabbitmq可以将消息从队列中删除
		false,              // exclusive, 排他性，如果为true，只能被创建者使用，且连接断开时自动删除
		false,              // no-local, 如果为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,              // no-wait，不等待，如果为true，不等待服务器响应，直接返回
		nil,                // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}
	stop := make(chan bool)
	go func() {
		for {
			select {
			case d := <-msgs:
				// log.Printf("Received a message: %s", d.Body)
				consumeFunc(string(d.Body))
			case <-stop:
				return
			}
		}
	}()
	<-stop
}

// Close 关闭channel和connection。注意：先关闭channel，再关闭connection
func (mq *RabbitMQ) Close() {
	mq.channel.Close()
	mq.conn.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

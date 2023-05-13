package rabbitmq

import (
	"log"

	// "github.com/streadway/amqp" // 不再维护
	amqp "github.com/rabbitmq/amqp091-go"
)

type client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// rabbitmq url, 格式: amqp://user:pass@host:port/vhost
	url      string
	qname    string
	exchange string
	key      string
}

func newClient(url, qname, exchange, key string) *client {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return &client{
		conn:     conn,
		channel:  channel,
		url:      url,
		qname:    qname,
		exchange: exchange,
		key:      key,
	}
}

// Close 关闭channel和connection。注意：先关闭channel，再关闭connection
func (mq *client) Close() {
	mq.channel.Close()
	mq.conn.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

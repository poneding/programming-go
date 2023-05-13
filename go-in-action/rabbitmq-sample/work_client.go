package rabbitmq

// Work 模式，一个生产者，多个消费者，消息只能被一个消费者消费
type WorkClient struct {
	*SimpleClient
}

func NewWorkClient(url, qname string) *WorkClient {
	return &WorkClient{
		NewSimpleClient(url, qname),
	}
}

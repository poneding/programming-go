package load_balancer_demo

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	LoadBalancer struct {
		Clients []*Client
		Size    int
	}
	Client struct {
		ID string
	}
)

var x int

func (lb *LoadBalancer) GetClientByPolling() *Client {
	clients := lb.Clients
	i := x % (lb.Size)
	x++
	return clients[i]
}

func (lb *LoadBalancer) GetClientByRandom() *Client {
	clients := lb.Clients

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond)
	r := rand.Intn(lb.Size)
	return clients[r%(lb.Size)]
}

func (lb *LoadBalancer) RegisterClient(clients ...*Client) {
	lb.Clients = append(lb.Clients, clients...)
	lb.Size += len(clients)
}

func (c *Client) Do() {
	fmt.Printf("Client %s doing", c.ID)
}

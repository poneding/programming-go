package load_balancer_demo

import "testing"

func Test(t *testing.T) {
	lb := &LoadBalancer{
		Clients: make([]*Client, 0),
	}
	lb.RegisterClient(&Client{"C1"}, &Client{"C2"}, &Client{"C3"})

	c := lb.GetClientByRandom()
	c.Do()

	c2 := lb.GetClientByRandom()
	c2.Do()
}

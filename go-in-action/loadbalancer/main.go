package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	lb := &LoadBalancer{
		virtualEndpoints: make(map[int]*Endpoint, 0),
	}
	lb.RegisterEndpoint(&Endpoint{
		ip:     "127.0.0.1",
		port:   8081,
		weight: 1,
	}, &Endpoint{
		ip:     "127.0.0.1",
		port:   8082,
		weight: 6,
	}, &Endpoint{
		ip:     "127.0.0.1",
		port:   8083,
		weight: 3,
	})

	fmt.Println("Resolve endpoint with round robin:")
	for i := 0; i < 10; i++ {
		lb.ResolveWithRoundRobin().Show()
	}

	fmt.Println("Resolve endpoint with random:")
	for i := 0; i < 10; i++ {
		lb.ResolveWithRandom().Show()
	}

	fmt.Println("Resolve endpoint with weight:")
	var resolveResult = make(map[int32]int, len(lb.endpoints))
	for i := 0; i < 10; i++ {
		ep := lb.ResolveWithWeight()
		resolveResult[ep.port]++
		ep.Show()
	}
	fmt.Println(resolveResult)

}

type (
	LoadBalancer struct {
		endpoints []*Endpoint

		virtualEndpoints map[int]*Endpoint
	}

	Endpoint struct {
		ip     string
		port   int32
		weight int32
	}
)

var x int

func (lb *LoadBalancer) ResolveWithRoundRobin() *Endpoint {
	i := x % len(lb.endpoints)
	x++
	return lb.endpoints[i]
}

func (lb *LoadBalancer) ResolveWithRandom() *Endpoint {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return lb.endpoints[r.Intn(len(lb.endpoints))%(len(lb.endpoints))]
}

func (lb *LoadBalancer) ResolveWithWeight() *Endpoint {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return lb.virtualEndpoints[r.Intn(len(lb.virtualEndpoints))%(len(lb.virtualEndpoints))]
}

func (lb *LoadBalancer) RegisterEndpoint(endpoints ...*Endpoint) {
	index := len(lb.virtualEndpoints)
	for _, v := range endpoints {
		for i := 0; i < int(v.weight); i++ {
			lb.virtualEndpoints[index] = v
			index++
		}
	}
	lb.endpoints = append(lb.endpoints, endpoints...)
}

func (c *Endpoint) Show() {
	fmt.Printf("Endpoint %s:%d\n", c.ip, c.port)
}

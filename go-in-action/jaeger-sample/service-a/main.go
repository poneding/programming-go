package main

import (
	"fmt"
	"io"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Service A: Calling Service B")

	req, err := http.NewRequest("GET", "http://service_a_envoy:8788/", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, "Service B: "+string(body))

	req, err = http.NewRequest("GET", "http://service_a_envoy:8791/", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, string(body))
}

// 调用关系 A => B => C
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

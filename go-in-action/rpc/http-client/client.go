package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	json := `{"method":"HelloService.Hello","params":["Jay Chou"],"id":1}`
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:1000/json-rpc", bytes.NewBuffer([]byte(json)))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	fmt.Printf("StatusCode: %s\n", string(body))
}

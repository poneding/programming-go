package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

func GetApi() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	defer resp.Body.Close()
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	fmt.Printf("StatusCode: %s\n", string(body))
}

func PostApi1() {
	json := `{"id":"u-001","name":"Jay","age":18}`
	resp, err := http.Post("https://example.com/user", "application/json", bytes.NewBuffer([]byte(json)))
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	fmt.Printf("StatusCode: %s\n", string(body))
}

func PostApi2() {
	// 1. json
	//json := `{"id":"u-001","name":"Jay","age":18}`
	//req, _ := http.NewRequest(http.MethodPost, "https://example.com/user", bytes.NewBuffer([]byte(json)))

	// 2. map
	reqBody, _ := json.Marshal(map[string]string{
		"id":   "u-001",
		"name": "Jay",
		"age":  "18",
	})
	req, _ := http.NewRequest(http.MethodPost, "https://example.com/user", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("My_Custom_Header", "Value")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	fmt.Printf("StatusCode: %s\n", string(body))
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PostApi3() {
	user := User{
		Id:   "u-001",
		Name: "Jay",
		Age:  18,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(user)
	resp, err := http.Post("https://example.com/user", "application/json", buf)
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	fmt.Printf("StatusCode: %s\n", string(body))
}

// go get golang.org/x/oauth2
func getAccessToken() string {
	var authCfg = &clientcredentials.Config{
		ClientID:     "xxx",
		ClientSecret: "xxx",
		TokenURL:     "https://xxx.com/connect/token",
		EndpointParams: url.Values{
			"grant_type": {"client_credentials"},
		},
	}
	token, err := authCfg.TokenSource(context.Background()).Token()
	if err != nil {
		fmt.Errorf("get access token failed. ERROR: %s\n", err.Error())
	}
	return token.AccessToken
}

func OAuth2Api() {
	json := `{"id":"u-001","name":"Jay","age":18}`
	req, _ := http.NewRequest(http.MethodPost, "https://example.com/user", bytes.NewBuffer([]byte(json)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Bearer", getAccessToken())

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	fmt.Printf("StatusCode: %s\n", string(body))
}

func FileApi() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}

	defer file.Close()

	var reqBody bytes.Buffer
	multiPartWriter := multipart.NewWriter(&reqBody)
	fileWriter, err := multiPartWriter.CreateFormFile("file_field", "hello.txt")
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}

	fieldWriter, err := multiPartWriter.CreateFormField("normal_field")
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}

	_, err = fieldWriter.Write([]byte("value"))
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}

	multiPartWriter.Close()

	req, err := http.NewRequest(http.MethodPost, "http://example.com/file", &reqBody)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	fmt.Printf("StatusCode: %s\n", string(body))
}

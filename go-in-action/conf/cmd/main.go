package main

import (
	"fmt"

	"github.com/poneding/learning-go/practice/conf"
)

func main() {
	conf.SetConfigPath("./config.json")
	// conf.SetConfigPath("./config.yaml")
	c, err := conf.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}

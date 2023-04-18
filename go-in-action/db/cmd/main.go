package main

import (
	"fmt"

	"github.com/poneding/learning-go/practice/db/cmd/repo"
)

func main() {
	repo.DB.Create(&repo.User{
		ID:       "u-001",
		Name:     "dp",
		Password: "123456",
		IsAdmin:  true,
	})

	var u repo.User
	repo.DB.Get("u-001", &u)
	fmt.Println(u)
}

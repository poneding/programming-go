package main

import (
	"fmt"
	"log"
	"os"

	"goemail"
)

func main() {
	emailHost := "smtp-mail.outlook.com"
	var emailPort int32 = 587
	emailUser := "poneding@outlook.com"
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	client, err := goemail.InitEmailClient(emailHost, emailPort, emailUser, emailPassword)
	if err != nil {
		log.Fatalln(err)
	}

	emailModel := goemail.EmailModel{
		Subject: "Test Email",
		From:    "poneding@outlook.com",
		To:      []string{"poneding@gmail.com"},
		Cc:      []string{"pding@mail.com"},
		Body:    "This is an email from golang test client.",
	}
	err = client.SendEmail(emailModel)
	if err != nil {
		fmt.Println("send email err:", err)
		return
	}
	fmt.Println("please check mailbox")
}

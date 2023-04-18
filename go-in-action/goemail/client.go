package goemail

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

type EmailClient struct {
	Host     string
	Port     int32
	User     string
	Password string
}

func InitEmailClient(host string, port int32, user, password string) (*EmailClient, error) {
	if len(host) == 0 || len(user) == 0 || len(password) == 0 || port == 0 {
		return nil, errors.New("invalid args")
	}
	emailClient := &EmailClient{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}
	return emailClient, nil
}

func (c *EmailClient) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (c *EmailClient) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(c.User), nil
		case "Password:":
			return []byte(c.Password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

type EmailModel struct {
	Subject string
	From    string
	To      []string
	Cc      []string
	Body    string
}

func (client *EmailClient) SendEmail(email EmailModel) error {
	if client == nil {
		return errors.New("email client is not initialized")
	}
	c, err := smtp.Dial(fmt.Sprintf("%s:%d", client.Host, client.Port))

	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: client.Host, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			fmt.Println("call start tls")
			return err
		}
	}

	if ok, _ := c.Extension("AUTH"); ok {
		if err = c.Auth(client); err != nil {
			fmt.Println("check auth with err:", err)
			return err
		}
	}

	if err = c.Mail(email.From); err != nil {
		return err
	}
	for _, addr := range email.To {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}

	header := make(map[string]string)
	header["Subject"] = email.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = `text/plain; charset="utf-8"`
	header["Content-Transfer-Encoding"] = "base64"
	header["To"] = strings.Join(email.To, ",")
	if len(email.Cc) > 0 {
		header["Cc"] = strings.Join(email.Cc, ",")
	}
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(email.Body))
	_, err = w.Write([]byte(message))

	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

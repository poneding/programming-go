package main

import (
	"gorm-viper/conf"
	"log"
)

type User struct {
	Id       string
	UserName string `mapstructure:"name"`
}

func main() {
	readFromJsonFile()
	readFromYamlFile()
	readFromYmlFile()
	readFromTomlFile()
	readFromIniFile()
	readFromEnvFile()
	readFromDestFile()
}

func readFromJsonFile() {
	c := conf.ReadFrom("./sample/conf/conf.json")
	log.Println(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Println(u.Id)
	log.Println(u.UserName)
}

func readFromYamlFile() {
	c := conf.ReadFrom("./sample/conf/conf.yaml")
	log.Println(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Println(u.Id)
	log.Println(u.UserName)
}

func readFromYmlFile() {
	c := conf.ReadFrom("./sample/conf/conf.yml")
	log.Println(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Println(u.Id)
	log.Println(u.UserName)
}

func readFromTomlFile() {
	c := conf.ReadFrom("./sample/conf/conf.toml")
	log.Println(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Println(u.Id)
	log.Println(u.UserName)
}

func readFromIniFile() {
	c := conf.ReadFrom("./sample/conf/conf.ini")

	log.Println(c.Get("default.key1"))
	log.Println(c.Get("user.id"))
	log.Println(c.Get("user.name"))
}

func readFromEnvFile() {
	c := conf.ReadFrom("./sample/conf/conf.env")

	log.Println(c.Get("KEY_1"))
	log.Println(c.Get("USER_ID"))
	log.Println(c.Get("USER_NAME"))
}

func readFromDestFile() {
	c := conf.ReadFromOptions(conf.ConfigOptions{
		Path: "./sample/conf",
		Type: "json",
		Name: "app.conf",
	})
	log.Println(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Println(u.Id)
	log.Println(u.UserName)
}

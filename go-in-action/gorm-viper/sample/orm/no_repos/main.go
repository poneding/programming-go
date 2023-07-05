package main

import (
	"fmt"
	"gorm-viper/orm"
	"gorm-viper/sample/orm/no_repos/entities"
	"time"

	"log"
)

var r *orm.Repository

func int() {
	r = &orm.Repository{
		Database: orm.NewDatabase(orm.DatabaseConfig{
			Type:     "mysql",
			Username: "admin",
			Password: "123456",
			Database: "test",
			Host:     "127.0.0.1",
			Port:     3306,
		}),
	}
}

func main() {
	CreateIsExistGetFirstDelete()
	List()
	Page()
	All()
}

func CreateIsExistGetFirstDelete() {
	// Create
	create := &entities.User{
		Name: "Jay Chou",
		Age:  time.Now().Year() - 1979,
	}
	r.Create(&create)

	// IsExist
	isExist1 := r.IsExist(create.Id, &entities.User{})
	isExist2 := r.IsExist("not-exist-id", &entities.User{})
	log.Printf("isExist1: %T", isExist1)
	log.Printf("isExist2: %T", isExist2)

	// Get
	var get entities.User
	r.Get(create.Id, &get)
	log.Printf("get.Name: %s", get.Name)

	// First
	var first entities.User
	r.First(&get, orm.FirstOption{
		Condition: orm.Condition("name = ?", "Jay Chou"),
	})
	log.Printf("first.Name: %s", first.Name)

	// Delete
	var delete entities.User
	r.Delete(create.Id, &delete)
	log.Printf("delete.Name: %s", delete.Name)
}

func List() {
	var users []entities.User
	for i := 0; i < 10; i++ {
		users = append(users, entities.User{
			Name: fmt.Sprintf("user-list-%d", i),
			Age:  i,
		})
	}
	r.Database.Create(&users)

	var list []entities.User
	r.List(&list, orm.ListOption{
		Condition: orm.Condition("name like ?", "%user-list-%"),
	})

	log.Println("list users:")
	for _, u := range list {
		log.Println(u.Name)
	}

	r.Sweep(&entities.User{}, orm.SweepOption{
		Condition: orm.Condition("name like ?", "%user-list-%"),
	})
}

func Page() {
	var users []entities.User
	for i := 0; i < 30; i++ {
		users = append(users, entities.User{
			Name: fmt.Sprintf("user-page-%d", i),
			Age:  i,
		})
	}
	r.Database.Create(&users)

	var page []entities.User
	r.Page(&page, orm.PageOption{
		Condition: orm.Condition("name like ?", "%user-page-%"),
		No:        2,
		Size:      10,
	})

	log.Println("page users:")
	for _, u := range page {
		log.Printf("u.ip: \t%s,\tu.Name: \t%s\t", u.Id, u.Name)
	}

	r.Sweep(&entities.User{}, orm.SweepOption{
		Condition: orm.Condition("name like ?", "%user-page-%"),
	})
}

func All() {
	var users []entities.User
	for i := 0; i < 5; i++ {
		users = append(users, entities.User{
			Name: fmt.Sprintf("user-all-%d", i),
			Age:  i,
		})
	}
	r.Database.Create(&users)

	var all []entities.User
	r.All(&all)
	log.Println("all users:")
	for _, u := range all {
		log.Printf("u.ip: \t%s,\tu.Name: \t%s\t", u.Id, u.Name)
	}

	r.Sweep(&entities.User{}, orm.SweepOption{
		Condition: orm.Condition("name like ?", "%user-all-%"),
	})
}

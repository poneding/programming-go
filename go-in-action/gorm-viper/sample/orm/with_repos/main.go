package main

import (
	"fmt"
	"gorm-viper/orm"
	"gorm-viper/sample/orm/with_repos/entities"
	"gorm-viper/sample/orm/with_repos/repos"
	"log"
	"time"
)

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
	repos.User().Create(&create)

	// IsExist
	isExist1 := repos.User().IsExist(create.Id, &entities.User{})
	isExist2 := repos.User().IsExist("not-exist-id", &entities.User{})
	log.Printf("isExist1: %T", isExist1)
	log.Printf("isExist2: %T", isExist2)

	// Get
	get, _ := repos.User().Get(create.Id) // Get method is override.
	log.Printf("get.Name: %s", get.Name)

	// GetByName
	getByName, _ := repos.User().GetByName("Jay Chou")
	log.Printf("getByName.Name: %s", getByName.Name)

	// First
	var first entities.User
	repos.User().First(&get, orm.FirstOption{
		Condition: orm.Condition("name = ?", "Jay Chou"),
	})
	log.Printf("first.Name: %s", first.Name)

	// Delete
	var delete entities.User
	repos.User().Delete(create.Id, &delete)
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
	repos.User().Database.Create(&users)

	var list []entities.User
	repos.User().List(&list, orm.ListOption{
		Condition: orm.Condition("name like ?", "%user-list-%"),
	})

	log.Println("list users:")
	for _, u := range list {
		log.Println(u.Name)
	}

	repos.User().Sweep(&entities.User{}, orm.SweepOption{
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
	repos.User().Database.Create(&users)

	var page []entities.User
	repos.User().Page(&page, orm.PageOption{
		Condition: orm.Condition("name like ?", "%user-page-%"),
		No:        2,
		Size:      10,
	})

	log.Println("page users:")
	for _, u := range page {
		log.Println("u.ip: \t%s,\tu.Name: \t%s\t", u.Id, u.Name)
	}

	repos.User().Sweep(&entities.User{}, orm.SweepOption{
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
	repos.User().Database.Create(&users)

	var all []entities.User
	repos.User().All(&all)
	log.Println("all users:")
	for _, u := range all {
		log.Printf("u.ip: \t%s,\tu.Name: \t%s\t", u.Id, u.Name)
	}

	repos.User().Sweep(&entities.User{}, orm.SweepOption{
		Condition: orm.Condition("name like ?", "%user-all-%"),
	})
}

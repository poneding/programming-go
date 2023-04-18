package main

import "fmt"

// func main() {
// 	pase_user()
// }

type User struct {
	Name string
	Age  int
}

func pase_user() {
	m := make(map[string]*User)
	users := []User{
		{
			Name: "u1",
			Age:  18,
		},
		{
			Name: "u2",
			Age:  19,
		},
		{
			Name: "u3",
			Age:  20,
		},
	}

	for _, u := range users {
		ptr_u := u
		m[u.Name] = &ptr_u
	}

	//golang 的  for ... range 语法中， stu 变量会被复⽤，每次循环会将集合中的值复制
	//给这个变量，因此，会导致最后 m 中的 map 中储存的都是 stus 最后⼀个student的值。
	for _, v := range m {
		fmt.Print(v.Name)
	}
	// fmt.Print(m)
}

package main

import "fmt"

func main4() {
	var data = []int{1, 2, 3}
	for _, d := range data {
		d *= d
	}
	fmt.Println(data)

	for i, d := range data {
		data[i] = d * d
	}
	fmt.Println(data)

	var data2 = []string{"1", "2", "3"}
	for _, d := range data2 {
		d += d
	}
	fmt.Println(data2)

	for i, d := range data2 {
		data2[i] = d + d
	}
	fmt.Println(data2)

	var data3 = []User{
		{
			Name: "Pone",
			Age:  28,
		}, {
			Name: "Jay",
			Age:  40,
		},
	}
	for _, d := range data3 {
		d.Name += d.Name
		d.Age += d.Age
	}
	fmt.Println(data3)

	for i, d := range data3 {
		data3[i].Name = d.Name + d.Name
		data3[i].Age = d.Age + d.Age
	}
	fmt.Println(data3)
}

type User struct {
	Name string
	Age  int
}

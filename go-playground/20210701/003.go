package main

// func main() {
// 	var a interface{}
// 	if a == nil {
// 		fmt.Println("a is nil")
// 	}
// 	if live() == nil {
// 		fmt.Println("A")
// 	} else {
// 		fmt.Println("B")
// 	}
// }

type People interface {
	Show()
}

type Student struct{}

func (s *Student) Show() {

}

func live() People {
	var s *Student
	return s
}

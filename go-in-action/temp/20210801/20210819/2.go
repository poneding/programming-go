package main

//import "fmt"
//
//// goconvey
//
//func main() {
//	//select 机制用来处理异步 IO 问题
//	//	select 机制最大的一条限制就是每个 case 语句里必须是一个 IO 操作
//
//	//var a interface{}
//	//a = "12"
//	//
//	//switch a.(type) {
//	//case int:
//	//	fmt.Println("int", a.(int))
//	//case string:
//	//	fmt.Println("string", a.(string))
//	//}
//	//
//	//b, ok := a.(string)
//	//if ok {
//	//	fmt.Println("ok", b)
//	//} else {
//	//	fmt.Println("no")
//	//}
//
//	s := slice([]int{1, 2, 3})
//	s.Add1(4)
//	fmt.Println(s) // 不会新增4
//	s.Add2(5)
//	fmt.Println(s) // 会新增5
//
//	Add3(s, 6)
//	fmt.Println(s) // 不会新增6
//	Add4(&s, 7)
//	fmt.Println(s) // 会新增7
//}
//
//type slice []int
//
//func (s slice) Add1(i int) {
//	s = append(s, i)
//}
//
//func (s *slice) Add2(i int) {
//	*s = append(*s, i)
//}
//
//func Add3(s slice, i int) {
//	s = append(s, i)
//}
//
//func Add4(s *slice, i int) {
//	*s = append(*s, i)
//}

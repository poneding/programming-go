package main

//func main() {
//	var data *byte
//	var in interface{}
//
//	fmt.Println(data, data == nil) // <nil> true
//	fmt.Println(in, in == nil) // <nil> true
//
//	in = data
//	fmt.Println(in, in == nil) // <nil> false // data 值为 nil，但 in 值不为 nil
//}

//func main() {
//	//interface 类型的变量只有在类型和值均为 nil 时才为 nil
//	// 如果知道类型了，那么不是nil
//
//	doIt := func(arg int) interface{} {
//		var result *struct{} = nil
//
//		if arg > 0 {
//			result = &struct{}{}
//		} else {
//			return nil // 明确指明返回 nil
//			//return result //返回的interface不是nil，但是对应的值是nil
//		}
//
//		return result
//	}
//
//	if res := doIt(-1); res != nil {
//		fmt.Println("Good result: ", res)
//	} else {
//		fmt.Println("Bad result: ", res) // Bad result: <nil>
//	}
//}

package reflect_demo

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `j:"nameaaa"`
}

func (a *User) Hi(to string) {
	fmt.Printf("Hi %s, my name is %s", to, a.Name)
}

func Do() {
	a := &User{}
	v := reflect.ValueOf(a) //将返回一直指针对象
	t := reflect.TypeOf(*a)
	e := v.Elem() //通过反射获取指针指向的元素类型，等效于对指针类型变量做了一个*操作

	f := e.FieldByName("Name")
	f.Set(reflect.ValueOf("Pone Ding"))
	//f.SetString("PoneDing")

	fName, ok := t.FieldByName("Name")
	if ok {
		fmt.Println("tag:", fName.Tag.Get("j"))
	}

	m := v.MethodByName("Hi")
	m.Call([]reflect.Value{reflect.ValueOf("Jay")})
}

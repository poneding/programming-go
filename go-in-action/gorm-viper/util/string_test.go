package util

import (
	"testing"
)

type U struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func TestJsonMarshal(t *testing.T) {
	v := U{
		Name: "jay",
		Age:  41,
	}
	t.Log(JsonMarshal(v))
}

func TestJsonUnmarshal(t *testing.T) {
	var u U
	JsonUnmarshal(`{"name":"jay","age":41}`, &u)
	t.Log(u.Name)
	t.Log(u.Age)
}

func TestCurrentFuncName(t *testing.T) {
	t.Log(CurrentFuncName())
}

func TestCurrentFuncFullName(t *testing.T) {
	t.Log(CurrentFuncFullName())
}

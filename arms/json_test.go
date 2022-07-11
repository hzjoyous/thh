package arms

import (
	"fmt"
	"testing"
)

func TestJsonEncode(t *testing.T) {
	type tmp struct {
		Name string
	}
	fmt.Println(JsonEncode(tmp{Name: "name"}))
}

func TestJsonDecode(t *testing.T) {
	type tmp struct {
		Name string
	}
	fmt.Println(JsonDecode[tmp](`{"name":"name"}`))
}

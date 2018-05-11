package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}

	st := reflect.TypeOf(s)

	field := st.Field(0)

	fmt.Println("color:\t", field.Tag.Get("color"))
	fmt.Println("species:\t", field.Tag.Get("species"))

}

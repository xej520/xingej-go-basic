package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string "user name" //引号里面的内容，就是tag
	Passwd string "user password"
}

func main() {

	user := &User{"chronos", "123456"}

	s := reflect.TypeOf(user).Elem()

	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag)
	}

}

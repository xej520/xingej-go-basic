//如果通过反射，对复杂类型进行修改
package main

import (
	"reflect"
	"fmt"
)

type Teac struct {
	Id int
	Name string
	Age int
}

func main() {
	teac := Teac{Id:5,Name:"Ant-man",Age:23}
	fmt.Println("teac:\t", teac)
	//传递的是 地址哦
	Set(&teac)
	fmt.Println("teac:\t", teac)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Printf("xxx")
		return
	}else{
		v = v.Elem()
	}

	// 通过FieldByName 这个方法，直接输入 名称，来获取
	f := v.FieldByName("Name")

	//校验，是否取到Name属性的值
	if !f.IsValid() {
		fmt.Printf("BAD")
		return
	}

	//然后，再校验，类型是否匹配
	if f.Kind() == reflect.String {
		f.SetString("Iron Man")
	}




}






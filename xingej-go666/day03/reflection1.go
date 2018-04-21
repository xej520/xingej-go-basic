//反射  练习
//获取字段的类型信息，方法信息， 属性信息
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (user User) Hello(){
	fmt.Println("hello world")
}

func main() {
	user := User{Id:2, Name:"xiaoqiang",Age:30}
	info(user)
}

//定义一个方法
//此方法接收的是一个空接口
//这样的话，其实是任何类型都可以接收了
func info(o interface{}) {
	//获取o的类型
	t := reflect.TypeOf(o)

	fmt.Println("Type:\t", t.Name())

	//取出类型后，需要对类型进行校验，
	//查看是否是规定的类型
	if k :=t.Kind(); k != reflect.Struct {
		//如果不是规定的Struct类型的话，就需要进行
		//异常处理
		fmt.Println("输入参数的类型不匹配!")
		return
	}

	//获取o的属性
	v := reflect.ValueOf(o)
	fmt.Println("Fields:\n")

	//开始遍历，输出字段
	for i := 0; i < t.NumField(); i++ {
		//获取属性下标
		f := t.Field(i)
		val := v.Field(i).Interface()
		//输出属性的名称，类型，对应的值
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val)
	}

	//开始获取 方法的基本信息
	for i:=0; i<t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}

}










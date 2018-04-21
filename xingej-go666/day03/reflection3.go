//如果通过反射，对基本类型进行修改
package main

import (
	"reflect"
	"fmt"
)

func main() {
	//下面测试，对基本类型的修改 操作
	x := 456
	//传递的参数是  地址
	v := reflect.ValueOf(&x)

	//Elem方法，是取出元素值来，然后通过setint方法，进行重新设置
	v.Elem().SetInt(789)

	fmt.Println(x)
}





//defer 测试
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//int类型
	i := 9
	//转换成string类型
	num := strconv.Itoa(i)

	//reflect.TypeOf(num).Name()获取类型名字
	if reflect.TypeOf(num).Name() == "string" {
		fmt.Println("==1==>:\t", num)
	} else {
		fmt.Println("==2==>:\t", num)
	}

}

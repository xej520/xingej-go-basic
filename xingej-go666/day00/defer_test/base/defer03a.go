package main

import "fmt"

func main() {

	for i:=0; i<100; i++{
		//这里defer调用的不是匿名函数，所以i是局部变量
		defer fmt.Println(i)

		if i == 30 {
			panic("println too many")
		}
	}

}

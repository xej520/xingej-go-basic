//匿名函数  练习
package main

import "fmt"

func main() {

	// 匿名函数，只在一处使用
	// 并且没有名字 ，所以，才称为，匿名函数
	// 并且 将匿名函数，赋值给了一个变量
	nimingFunction := func() {
		fmt.Println("this is a niming function")
	}

	//调用匿名函数
	nimingFunction()
}



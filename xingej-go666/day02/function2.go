//函数 练习
package main

import "fmt"

func main() {
	a := 5
	A(a)

	fmt.Println("main方法中打印:\t", a)  // 5
}

//值  传递 ，而非 地址的传递
//这种方式的修改，不会影响到外面的
func A(a int) {
	a = 3
	fmt.Println("函数A内部打印:\t", a) //3
}




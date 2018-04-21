// recover 测试
// recover 要在defer 中使用
package main

import "fmt"

func main() {
	C()
	C02()
	C03()
}

func C() {
	fmt.Println("this is C function")
}

func C02()  {
	panic("panic is  C02")
}

func C03()  {
	fmt.Println("this is C03 function")
}





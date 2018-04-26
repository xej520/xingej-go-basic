//defer 测试
package main

import "fmt"

func main() {
	defer fmt.Println(" A")
	defer fmt.Println(" B")
	defer fmt.Println(" C")
	defer fmt.Println(" D")

	//defer 调用匿名函数
	defer func() {
		fmt.Println(" E")
	}()

	//defer调用普通函数
	defer deferTest()

}

func deferTest() {
	fmt.Println("====deferTest====")
}

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
	//fmt.Println("=======================")
	//for  i := 0; i<3; i++{
	//	defer fmt.Println(i)
	//}

	//执行时，会将所有的defer 倒着执行的， 如输出结果，如下所示：
	//2
	//1
	//0
	//D
	//C
	//B
	//A
}

func deferTest() {
	fmt.Println("====deferTest====")
}

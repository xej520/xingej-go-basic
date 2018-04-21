package main

import "fmt"

func main() {
	b1()
	b2()
	b3()
}

func b1() {
	fmt.Println("b1")
}
func b2() {

	//defer 定义的匿名函数，一定要在painc前面
	defer func() {

		if error := recover(); error != nil {
			fmt.Println("Recover in b2")
		}

	}()

	panic("b2 problem")
}

func b3() {
	fmt.Println("b3")
}

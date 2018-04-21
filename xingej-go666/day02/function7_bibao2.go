//闭包练习 2
package main

import "fmt"

func main() {
	//声明一个匿名函数，返回值也是个匿名函数
	add := func (arg int) func(int) int{
 		return func(i int) int {
			return arg + i
		}
 	}

 	test := add(6)
 	test2 := add(6)

 	fmt.Println("result:\t", test(10))
 	fmt.Println("result:\t", test2(10))
 }



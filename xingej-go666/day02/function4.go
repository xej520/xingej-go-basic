//函数 参数 ，传递的是地址
package main

import "fmt"

func main() {

	a := 6
	A1(&a)
	fmt.Println("main:\t", a)

}

// 传递的是地址， 也会改变 旧值的
func A1(a *int)  {
	//*a, 表示  取到 了值
	*a = 3
	fmt.Println("A:\t", *a)
}
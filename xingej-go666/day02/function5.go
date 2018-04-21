// 将函数 赋值给一个变量
//从此变量，就是函数了
package main

import "fmt"

func main() {
	//将B 函数，赋值给变量bias_for_B
	// 这就说明，Go语言中，一切皆类型
	bias_for_B := B

	//调用函数
	bias_for_B()
}

func B() {
	fmt.Println("this is B() function!")
}





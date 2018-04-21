//闭包函数  测试
//闭包函数 的唯一作用，就是
// 返回一个 匿名函数
package main

import "fmt"

func main() {
	//其中，fun 其实，就是一个函数类型变量
	fun := closure(10)

	fmt.Println(fun(2))
	fmt.Println(fun(4))

}

//返回的是 匿名函数
// 此 匿名函数，需要一个参数输入，并且返回一个int类型的值
func closure(x int) func(int) int {

	return func(i int) int {
		fmt.Printf("--> %d", i)
		fmt.Println()
		return x + i
	}
}

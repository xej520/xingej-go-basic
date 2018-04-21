//当前程序的包名，必须是非注释代码的第一行
package main

import "fmt"

//常量的定义
const PI = 3.14

const (
	const1 = "1"
	const2 = 2
	const3 = 3
)

//全局变量的声明与赋值

func main() {
	fmt.Println("const1:\t", const1)
	fmt.Println("const2:\t", const2)
	fmt.Println("const3:\t", const3)

	fmt.Println("---------------------")

	//字符串连接
	fmt.Println("const1:\t", const1, "\nconst2:\t", const2, "\nconst3:\t" , const3)
}

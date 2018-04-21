//channel 练习
package main

import "fmt"

func main() {

	//在内存里开辟了一个chan，通道
	//存储的是int类型
	c := make(chan int)

	//开始往通道里，存储数据
	c <- 2
	c <- 4

	//从通道里读取数据
	fmt.Println("===>:\t", <-c)
	fmt.Println("===>:\t", <-c)
}




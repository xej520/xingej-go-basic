package main

import "fmt"

func main() {

	c := make(chan int, 2)

	c <- 2
	c <- 3
	c <- 4

	//从通道里读取数据
	fmt.Println("===>:\t", <-c)
	fmt.Println("===>:\t", <-c)
	fmt.Println("===>:\t", <-c)

}


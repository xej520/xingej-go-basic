//channel 作为 返回类型，进行传递
package main

import "fmt"

func main() {
	//声明一个数组，类型是chan类型
	var channels [10]chan int

	//对这个数组，进行初始化
	for i:=0;i<10 ;i++  {
		channels[i] = createWorker(i)
	}

	//开始发送数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

func createWorker(i int) chan int {
	c := make(chan int)
	//创建chan后，必须在这里，添加接收的，不然，
	go func() {

		for  {
			//对channel进行的业务逻辑处理
			fmt.Printf("Worker %d received %c \n", i, <-c)
		}
	}()

	return c
}

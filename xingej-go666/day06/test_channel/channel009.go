package main

import (
	"fmt"
	"time"
)

func main() {

	var channel chan int
	channel = createWorkerV3a()

	//开始给channel发送数据
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4

	//等发送方已经将数据发送完毕了，可是接收方，还在阻塞着
	//如何告诉 接收方，我已经发送完毕了呢？
	time.Sleep(time.Second * 10)

}

func createWorkerV3a() chan int{
	c := make(chan int)

	go func() {
		for  {
			//如果接收方，没有接收到数据的话，就会一直在这里阻塞着
			value := <- c
			fmt.Println("接收的数据是:\t", value)
		}

	}()
	return c

}



package main

import (
	"fmt"
	"time"
)

func main() {

	var channel chan int
	channel = createWorkerV3()

	//开始给channel发送数据
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4

	time.Sleep(time.Second)

	//发送方，关闭channel
	//告诉，接收方，我这边已经发送完毕了
	close(channel)
}

func createWorkerV3() chan int{
	c := make(chan int)

	go func() {
		for  {
			value , ok := <- c

			if !ok {
				fmt.Println("发送方，已经发送完毕了!", value)
				break
			}

			fmt.Println("接收的数据是:\t", value)
		}

	}()
	return c

}



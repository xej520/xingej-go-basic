package main

import (
	"fmt"
	"time"
)

func main() {

	var channel chan int

	channel = createWorkerV4()

	channel <- 3
	channel <- 4
	channel <- 5
	channel <- 35

	close(channel)

	time.Sleep(time.Second)

}

func createWorkerV4() chan int {

	c := make(chan int)
	go func() {
		//如果，没有数据的话，这里会进行阻塞
		for v := range c {
			fmt.Println("====> 读取到的数据是；\t", v)
		}

		//这条语句，只有发送方，主动将channel关闭后，
		//才会打印的
		fmt.Println("=====接收方已经将数据发送完毕了=====")

	}()

	return c

}


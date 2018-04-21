// 测试chan 只能作为 接收chan类型使用
package main

import "fmt"

func main() {
	//声明 chan类型，只能用作读取
	var channels [10]<-chan int

	for i := 0; i < 10; i++ {
		channels[i] = createWorkerChannelReceive(i)
	}

	for i := 0; i < 10; i++ {
		//这里只能读取数据了
		fmt.Printf("channel[%d]  receive %c \n", i, <-channels[i])
	}

}

// 只能作为接收数据使用
func createWorkerChannelReceive(index int) <-chan int {
	//创建的是双向通道，只是，返回的时候，添加了限制
	c := make(chan int)
	go func() {
		for  {
			//不断的往这个chan里发送数据
			c <- 'a' + index
		}
	}()

	return c
}

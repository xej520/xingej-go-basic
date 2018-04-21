//在go语言中，channel也是一等公民，可以做为参数进行传递
package main

import "fmt"

func main() {
	var channnels [10]chan int
	for i := 0; i < 10; i++ {
		channnels[i] = make(chan int)
		go worker(channnels[i], i)
	}

	//开始发送数据
	for i := 0; i < 10; i++ {
		channnels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channnels[i] <- 'A' + i
	}
}

//测试，channel 作为参数进行传递
func worker(c chan int, index int) {
	for {
		fmt.Printf("worker %d receive %c\n", index, <-c)
	}
}

//测试，不建议在goroutines里使用fmt.println语句
//因为当有很有goroutines，同时都用到fmt.println的话
//要控制谁到底是谁在使用stdout，stdin
//一般使用channel通道，将信息通过channel通道发给主main来实现
package main

import "fmt"

func main() {
	msg := make(chan string)
	go func() {
		msg <- "hello, goroutines!"
	}()

	fmt.Println("收到的消息是:\t", <-msg)
}

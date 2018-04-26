package main

import (
	"fmt"
	"time"
)

//发送数据
func sample(msg chan string) {
	msg <- "hello goroutines 1"
	msg <- "hello goroutines 2"
	msg <- "hello goroutines 3"
	msg <- "hello goroutines 4"
}

//接收数据
func sample2(msg chan string) {
	time.Sleep(2 * time.Second)
	str := <-msg
	msg <- str + " hello world goroutines!"

	close(msg)
}

func main() {
	msg := make(chan string, 3)
	go sample(msg)
	go sample2(msg)
	time.Sleep(3 * time.Second)
	for key := range msg {
		fmt.Println(key)
	}

	//fmt.Println(<-msg)
	//fmt.Println(<-msg)
	//fmt.Println(<-msg)
	//fmt.Println(<-msg)

}

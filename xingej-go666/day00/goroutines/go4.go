package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool, 2)

	go func() {
		fmt.Println("-----")
		time.Sleep(time.Second * 5)
	}()

	//channel是缓存类型的情况下
	//如果channel里没有数据的话，
	//主线程也会在这里阻塞着
	//
	<-c

	fmt.Println("==================")
}

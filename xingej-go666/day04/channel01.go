package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	singleChannel := make(chan os.Signal, 1)
	dataChannel := make(chan int)
	for {
		select {
		case <-singleChannel:
			fmt.Println("要关闭程序了!")
			close(singleChannel)
			return
		}
	}

	go func() {
		for {
			dataChannel <- 3
		}
	}()

	go func() {

		for {
			fmt.Println("====>:\t", <-dataChannel)
		}
	}()

	time.Sleep(5 * time.Second)
}

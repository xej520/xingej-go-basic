//将select 用于发送数据  1
package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {

		for v := range c {
			fmt.Println("c =:\t", v)
		}

	}()

	for {
		select {
		//将select 用于 发送信息，
		case c <- 0:
		case c <- 1:
		}
	}
}

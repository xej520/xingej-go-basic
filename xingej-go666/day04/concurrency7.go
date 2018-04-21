//通道缓存 练习
package main

import "fmt"

func main() {
	//无缓存 类型的channel
	c := make(chan bool)

	go func() {
		fmt.Printf("hello ,Go world!")
		c <- true
	}()

	<-c

}




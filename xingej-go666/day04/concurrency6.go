//通道缓存 练习
package main

import "fmt"

func main() {
	// 有缓存 类型的channel，
	c := make(chan bool, 2)

	go func() {
		fmt.Printf("hello ,Go world!")
		<-c
	}()

	c <- true

}

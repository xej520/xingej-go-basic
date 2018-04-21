package main

import "fmt"

func main() {
	//无缓存类型channel
	c := make(chan bool)

	go func() {
		fmt.Println("go")
		<- c
	}()

	c <- true

}


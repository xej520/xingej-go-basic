package main

import "fmt"

func main() {

	c := make(chan bool)

	go func() {
		fmt.Println("go go")

		c <- true

	}()

	for v := range c {
		fmt.Println("=====> 我是主程序, " , v)
	}

}

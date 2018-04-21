package main

import "fmt"

func main() {
	c := make(chan bool, 2)

	go func() {
		fmt.Println("go")

		<- c
	}()

	c <- true

}

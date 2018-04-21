package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for  {
			value := <-c

			fmt.Println("====>:\t", value)
		}

	}()

	c <- 2
	c <- 3

}


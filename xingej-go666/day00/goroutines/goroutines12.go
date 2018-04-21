package main

import "fmt"

func main() {

	//声明两个channel
	c1, c2 := make(chan int), make(chan string)

	o := make(chan bool, 1)

	go func() {

		//监控channel
		for {
			select {
			case v1, ok := <-c1:
				if !ok {
					fmt.Println("===v1====")
					o <- true
					break
				}
				fmt.Println("===v1:\t", v1)

			case v2, ok := <-c2:
				if !ok {
					fmt.Println("====v2====")
					o <- true
					break
				}
				fmt.Println("===v2:\t", v2)
			}
		}

	}()

	c1 <- 1
	c2 <- "hello,"
	c1 <- 4
	c2 <- " Go world"

	close(c1)

	for i := 0; i < 2; i++ {
		<-o
	}

}

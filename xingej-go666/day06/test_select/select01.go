package main

import (
	"fmt"
)

func main() {
	//var c1, c2 chan int // 此时，c1, c2 都是 nil类型
	var c1, c2 = generator(), generator()

	worker := createWorker(0)

	n := 0

	hasValue := false

	for {
		var activeWorker chan<- int

		if hasValue {
			activeWorker = worker
		}

		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		default:
			fmt.Println("No value received!")
		}

	}

}

func generator() chan int {

	out := make(chan int)

	//利用goroutines发送数据
	go func() {
		i := 0
		for {
			out <- i
			i++
		}

	}()

	return out
}

func worker(i int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d \n", i, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)

	go worker(id, c)

	return c
}

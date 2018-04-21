package main

import (
	"sync"
	"fmt"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(10)

	for i :=0; i<10;i++{
		go sumV3(&wg, i)
	}

	wg.Wait()
}

func sumV3(sync *sync.WaitGroup, index int){
	sum := 0

	for i := 0; i<10000; i++{
		sum += i
	}

	fmt.Println(index, "\t", sum)

	sync.Done()

}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		fmt.Println("===go k8s====")

		time.Sleep(time.Second * 5)

		wg.Done()
	}()

	fmt.Println("----main---1----")
	//进行阻塞
	wg.Wait()

	fmt.Println("----main---2----")
}

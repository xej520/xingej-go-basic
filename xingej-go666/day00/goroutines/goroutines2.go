package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("go k8s")
	}()

	time.Sleep(time.Second * 2)
}

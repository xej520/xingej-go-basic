package main

import (
	"fmt"
)

func main() {
	flag := false

	go func() {
		fmt.Println("go k8s")
		flag = true
	}()

	//检测机制
	for {
		if flag {
			fmt.Println("====goroutines运行结束了======")
			break
		}
	}
}

package main

import "fmt"

func main() {

	defer fmt.Println(" world")
	defer fmt.Print(" k8s")

	//说明，return 之前定义的defer ，也会执行
	return

	defer fmt.Print(" hello")


}


package main

import "fmt"

func main() {

	defer fmt.Println(" world")
	defer fmt.Print(" k8s")

	// panic 之前定义的defer,都会执行的
	panic("error : clusterId is not exist!")

	defer fmt.Print(" hello")

}

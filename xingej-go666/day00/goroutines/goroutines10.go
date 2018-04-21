package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool)

	for i:=0; i<10; i++{
		go sumV2(c, i)
	}

	for i := 0; i<10; i++{
		<- c
	}

}

func sumV2(c chan bool, index int)  {
	sum := 0

	for i :=0; i<10000; i++{
		sum += i
	}

	fmt.Println(index, "\t", sum)

	c <- true

}
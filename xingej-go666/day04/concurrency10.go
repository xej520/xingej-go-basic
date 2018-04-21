package main

import (
	"fmt"
	"runtime"
)

func main() {
	//按照多核方式处理
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool)

	for i := 0; i<10 ;i++  {
		go showInfo(c, i)
	}

	for i :=0; i<10; i++{
		<-c
	}

}

func showInfo(c chan  bool, index int){
	a := 1
	for i:=0;i<10000000 ;i++  {
		a += i
	}

	fmt.Println(index,a )

	//
	c <- true
}





package main

import "fmt"

func main() {

	c := make(chan bool)

	for i:=0; i<10;i++ {
		go sum(c, i)
	}

	//读取
	<- c

}

func sum(c chan bool, index int) {
	sum := 0

	for i :=0; i<10000; i++{
		sum += i
	}

	fmt.Println(index,"\t" ,sum )

	if index == 9 {
		//如果传进来的是9的话，就往通道里存储存储true
		c <- true
	}

}


package main

import (
	"time"
	"fmt"
)

func main() {

	//sum := [10]int{}

	var sum [10]int

	for i :=0; i< 3; i++{
		go func(i int) {
			for{
				sum[i]++
			///fmt.Println("--->:\t", sum[i], "\t",i )
			}
		}(i)

	}

	//fmt.Println("==========可能马上退出了=======")
	time.Sleep(time.Second * 3)

	fmt.Println("==========main退出=======")

}


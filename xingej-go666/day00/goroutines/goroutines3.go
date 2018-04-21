package main

import "fmt"

func main() {
	//创建一个channel，channel的类型是bool类型
	ch1 := make(chan bool)

	go func() {
		fmt.Println("go k8s")
		//存储
		ch1 <- true
	}()

	//读取
	<-ch1

}

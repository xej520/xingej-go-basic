package main

import "fmt"

func main()  {
	//defer里，应该是有一个栈，先进后出
	defer fmt.Println(" world")
	defer fmt.Print(" k8s")
	defer fmt.Print(" hello")

	msgs := []string{" world", " k8s", "hello"}

	for _,v := range msgs {
		fmt.Print(v)
	}

}



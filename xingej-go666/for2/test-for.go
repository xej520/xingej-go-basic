package main

import "fmt"

func main() {

	flag := true

	// 如果flag 是true的话，就会一直运行下去，相当于 死循环的

	for flag {
		fmt.Println("=====")
		flag = false
	}

}

package main

import "fmt"

func main() {
	var num *int32
	*num = 5
	fmt.Println("=====>:\t", num)
	fmt.Println("=====>:\t", *num)
}

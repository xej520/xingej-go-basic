package main

import "fmt"

func adder() func(int) int {
	//可以简单的认为，
	//这是一个全局变量
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	first, end := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t %d ", first(i), end(-2*i))
		fmt.Println()
	}
}

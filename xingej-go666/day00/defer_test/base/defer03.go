package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		//defer，来调用 匿名函数
		//匿名函数时，i这里传递的是  就是 引用 ，
		defer func() {
			fmt.Printf("i = %d ", i)
			fmt.Println()
		}()

		if i == 30 {
			panic("println too many")
		}

	}
}

//defer3 中出现的问题，如何解决呢？
package main

import "fmt"

func main() {
	D01()
	D02()
	D03()
}

func D01()  {
	fmt.Println("this is D01")
}

//先写的是panic,然后，再写的是 defer
//这种写法是不行的
//因为，当程序运行到panic的时候，已经出错了，不会在往下执行了
//根本没有运行到defer
//func D02()  {
//	panic("this is D02")
//
//	defer func() {
//		if error := recover(); error != nil {
//			fmt.Println("Recover in D02")
//		}
//
//	}()
//
//}

// 因此，要将defer 写到panic前面，才行的
func D02()  {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Recover in D02")
		}

	}()

	panic("this is D02")
}

func D03()  {
	fmt.Println("this is D03")
}
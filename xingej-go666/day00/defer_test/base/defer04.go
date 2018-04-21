//为什么 会有 recover出现？
package main

import "fmt"

func c1() {
	fmt.Println("c1")
}

func c2() {
	panic("c1")
}

func c3() {
	fmt.Println("c1")
}

func main() {

	c1()
	c2()
	c3()

}


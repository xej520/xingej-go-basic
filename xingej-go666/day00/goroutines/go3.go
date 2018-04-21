package main

import "fmt"

func main() {

	ch := make(chan bool, 2)

	fmt.Println("===1==")

	if len(ch) == 0 {
		close(ch)
	}else {
		<-ch
	}

	fmt.Println("===2==")


}

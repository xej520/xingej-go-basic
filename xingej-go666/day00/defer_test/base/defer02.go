package main

import "fmt"

func main() {

	msgs := []string{" world", " k8s", "hello"}

	for _,v := range msgs {
		defer fmt.Print(v)
	}

}



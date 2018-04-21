package main

import "fmt"

type student struct {
	name string  `json:"id,omitempty"`
}

func main() {

	stu := student{"spark"}

	fmt.Println("===>:\t", stu)

}





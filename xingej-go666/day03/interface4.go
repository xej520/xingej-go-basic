package main

import "fmt"

func main() {

	test(spark{name:"laoxing", do: func() {
		fmt.Println("===============")
	}})

}

type spark struct {
	name string
	do func()
}

type hadoop struct {
	name string
	age int
}

func test(obj interface{}){
	//sparkType := obj.(hadoop)



	//fmt.Println("====>:\t", sparkType.do)

}



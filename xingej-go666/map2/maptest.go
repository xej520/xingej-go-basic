package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		fmt.Println("----1--->:\t")
		time.Sleep(1 * time.Second)
	}

}

func create() {

	recover()

	fmt.Println("-----------create-------")

	delete()

	update()

}

func delete() {
	fmt.Println("-----------create-------")
}

func update() {
	fmt.Println("-----------create-------")
}

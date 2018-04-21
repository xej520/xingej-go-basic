package main

import "fmt"

type annimal2 struct {
	name string

	show func()
}


func main() {
dog := annimal2{name:"xiaohuang",
		show: func() {
			fmt.Println("this is xiaohuang")
		},
	}

	dog.show()

	cat:= annimal2{name:"xiaohong",
		show: func() {
			fmt.Println("this is xiaohong")
		},
	}

	cat.show()

}




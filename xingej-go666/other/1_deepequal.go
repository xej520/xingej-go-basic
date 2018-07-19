package main

import (
	"fmt"
	"reflect"
)

func main() {

	oldNames := map[string]string{"school": "bjtu", "name": "spark"}
	newNames := map[string]string{"name": "spark", "school": "bjtu"}
	currentNames := map[string]string{"name": "spark", "school": "bjtu", "ID": "123456"}

	//跟顺序没有关系
	if reflect.DeepEqual(oldNames, newNames) {
		fmt.Println("----->: OK")
	}

	if !reflect.DeepEqual(oldNames, currentNames) {
		fmt.Println("----->: 不OK")
	}

}

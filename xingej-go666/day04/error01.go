package main

import (
	"errors"

	"fmt"
)

func a() (err error) {

	err = b(2000)

	if err != nil {
		return
	}

	fmt.Println("b2已经调用完成!")

	err = c2()
	if err != nil {
		return
	}

	err = errors.New("a内错误")

	return

}

func b(age int) (err error) {
	//具体的业务逻辑处理
	//如，参数校验
	if age < 1000 {
		fmt.Println("age:\t", age)
	} else {
		// 参数不符合要求
		err = errors.New("b内错误")
		fmt.Println("=========================")
	}

	return
}

func c2() (err error) {

	err = errors.New("c3333内错误")

	return

}

func main() {

	err := a()

	if err != nil {

		fmt.Println(err)

	}

}

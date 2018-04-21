//函数的修改
package main

import "fmt"

func main() {

	test001(3,4,2,1,4)

	fmt.Println("===========================")
	sl1 := []int{3,4,5}
	test02(sl1[0], sl1[1], sl1[2])
	//这种方式，是不会影响原来的值的
	fmt.Println(sl1)

	fmt.Println("===========================")
	test03(sl1)
	fmt.Println(sl1)

}

func test001(arg ... int)  {
	fmt.Println(arg)
}

//如果内部修改了变成参数的某个值，会不会影响到外面的值
// 说明，变长参数，是不会改变原来的值
//这种方式，本质上还是值的拷贝
func test02(arg ... int)  {
	arg[0] = 110
	arg[1] = 120
	fmt.Println(arg)
}

//若参数类型，直接是一个 切片类型 时，  会不会影响原值呢？
//经测试，这种方式，会修改原来的值
func test03(arg []int) {
	arg[0] = 110
	arg[1] = 120
	fmt.Println(arg)
}






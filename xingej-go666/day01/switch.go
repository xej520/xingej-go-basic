// switch 选择控制语句 练习
package main

import "fmt"

func main() {
	//
	num := 4
	//条件表达式的位置，
	// 1、写在switch后面
	//2、 写在每个case的后面
	// 不用显示的 写break，自动结束的
	switch num {
	case 1:
		fmt.Println("num=1")
	case 4:
		fmt.Println("num=4")
	default:
		fmt.Println("num=0")
	}

	fmt.Println("=====Switch语句的第2种表达形式=====")
	//将初始化 语句，写在switch后，记住，别忘记了  分号;
	// 但是，注意
	//age的作用域，仅仅是switch语句块，超出这个范围，就不行了
	// 因此，如果switch语句块结束后，还想使用的话，就不能在switch后面，初始化
	switch age := 4 ;{
	case age < 4:
		fmt.Println("age < 4")
	case age == 4:
		fmt.Println("age = 4")
		fallthrough
	case age >= 4:
		fmt.Println("age >= 4")
	default:
		fmt.Println("age = 0")
	}

	fmt.Println("=====Switch语句的第3种表达形式=====")
	//将条件表达式，写在case语句后面
	switch {
	case num < 4:
		fmt.Println("num < 4")
	case num == 4:
		fmt.Println("num = 4")
		//添加上fallthrough后，运行完这个case语句后，还会判断，剩下的case语句，是否满足条件，
		//如果满足条件，就继续运行
		fallthrough
	case num >= 4:
		fmt.Println("num >= 4")
	default:
		fmt.Println("num = 0")
	}

}

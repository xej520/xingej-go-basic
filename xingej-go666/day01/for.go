//for循环语句  练习测试
//注意，在Go语言中，只有for循环一个循环语句，没有while之类的
//但是，Go语言中的循环语句，功能比较强大，有
//3种形式
package main

import "fmt"

func main() {
	fmt.Println("======下面是第1种for循环形式======")
	var a = 1
	for {
		if a > 4{
			break
		}
		a++
		fmt.Println("a \t" , a)
	}

	fmt.Println("======下面是第2种for循环形式======")
	b:=4
	for b < 10 {
		fmt.Println("b:\t", b)
		b++
	}
	fmt.Println("======下面是第3种for循环形式======")
	for c:=1;c < 5 ;c++  {
		fmt.Println("c:\t", c)
	}


}


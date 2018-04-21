// 测试，defer 调用 匿名函数
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		// defer 调用匿名函数，因此，后面添加上了
		// 小括号()
		//并且，结合了闭包属性
		// i 是按照 地址传递的
		//最后一轮循环时，i已经变成了3了，因此
		//前两个的调用，也都改成了3了
		//而且defer是倒着执行的，
		defer func() {
			fmt.Println(i)
		}()
	}
}

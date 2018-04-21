//练习题
//运行以下程序，并分析输出结果，是否跟你想的一样的呢
package main

import "fmt"

func main() {

	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		//这里的i属于参数，按值传递的
		//defer定义，倒序执行
		defer fmt.Println("defer i = ", i)

		//属于闭包，引用的是地址
		defer func() {
			fmt.Println("defer_closure i = ", i)
		}()

		//属于闭包
		fs[i] = func() {
			//变量i是从外层循环中读过来的，
			//是引用地址
			fmt.Println("clousre i = ", i)
		}
	}

	for _, f := range fs {
		f()
	}

}

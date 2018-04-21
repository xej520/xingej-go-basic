//  并发  最基本的使用
// 使用go 关键字
package main

import (
	"fmt"
)

func main() {
	go Go()
}
//运行程序之后，为什么什么也没有输出呢？
//main函数在调用Go 函数时，需要创建启动一个goroutine，没有等待Go方法调用里面的
//语句，就返回退出了，当main退出之时，程序也就退出了
//也就是说，main函数，没有等待人家运行，就退出。因此，
//如何解决呢？
//最简单的方法，就是，在main方法里，运行go之后，开启一个sleep，等待一会，再结束


//声明一个function
func Go(){
	fmt.Printf("Go Go Go!!!")
}









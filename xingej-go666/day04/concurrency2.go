//使用最笨的方法，
package main

import (
	"fmt"
	"time"
)

func main() {
	//利用go关键字，开启一个goroutine
	go show()
	//main方法，先休息一下，等待上面的show方法运行完毕后，再结束

	// 也可以使用匿名函数
	go func() {
		fmt.Printf("======>hello go routine")
	}()

	time.Sleep(2 * time.Second)
}

func show()  {
	fmt.Printf("hello, GO!!!")
}


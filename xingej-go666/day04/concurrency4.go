//使用for range  以及 close（）来代替 <-
package main

import "fmt"

func main() {
	//要注意：
	//通道分为：单向通道，和双向通道
	//单向通道：只能取或者只能读
	//双向通道，都是可以的。
	//这里设置的是双向通道
	c := make(chan bool)

	go func() {
		fmt.Println("hello, Go world")

		// 你要是存储一次，c<-true， 下面的for v := range c 里c 就会运行两次
		//c <- true
		c <- true

		//必须明确的关闭c
		//这样，主程序main中for range 就不会阻塞等待了
		//close(c)
	}()

	for v := range c {
		fmt.Println("=====> 我是主程序, " , v)
	}


}




//select 练习
package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)

	//需要创建一个用于通信的信号灯
	//用于表示  c1，c2执行完毕了，
	//不然，主main线程，会马上就结束了
	o := make(chan bool, 1)

	go func() {
		//创建的是无线循环
		//如果不是无限循环的话，
		for {
			//select 运行一次，就往下执行了，不在接收了
			//java里也有类似的用法
			//如，SOCKET编程中，服务器端监听客户端
			//通过这种组合方式， 从而实现不断的信息接收与发送操作
			select {
			//将select用于读取数据信息， <- c1是读取的意思
			case v, ok := <-c1:
				//也就是说，当channel关闭的时候，就会退出select，就break
				if !ok {
					fmt.Println("====")
					o <- true
					break
				}
				fmt.Println("v1:\t", v)

			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("v2:\t", v)

			}
		}
	}()

	c1 <- 1
	c2 <- "hello,"
	c1 <- 4
	c2 <- " Go world"

	close(c1) //关闭c1时，select就会接收到，ok就会变成false,然后就会向channel里发送true，主线程读取到true
	//close(c2)

	//主程序读取
	for i := 0; i < 2; i++{
		<-o
	}


}

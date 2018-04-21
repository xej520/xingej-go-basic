//练习题：
//创建一个goroutine, 与主线程按顺序相互发送信息若干次，并打印
// 有点类似于 服务器与客户端交互的形式
package main

import "fmt"

var c chan string

func main() {

 	c = make(chan string)

 	go pingpang()

	for i := 0;i<10 ;i++  {
		//主线程，首先发送消息
		c <- fmt.Sprintf("From main: Hello, %d", i)
		//主线程，开始接收客户端发送过来的消息，并直接打印
		fmt.Println(<-c)
	}

}

func pingpang(){
	//客户端，死循环，不断的接收服务器发送过来的消息
	i := 0
	for{
		fmt.Println(<-c)
		//开始向服务器，发送信息
		c <- fmt.Sprintf("From client:Hi, %d", i)
		i++
	}
}

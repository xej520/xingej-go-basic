//通道缓存 练习
package main

import "fmt"

func main() {
	//无缓存 类型的channel
	c := make(chan bool)

	go func() {
		fmt.Printf("hello ,Go world!")
		// 读取c， 必须放在最后面，如果放在前面的话，就不能保证
		// 方法还没有运行完毕，主main已经结束了。
		<-c
	}()

	//
	c <- true
}

//运行结果，输出了内容:hello ,Go world!
//为什么会输出呢？
//channel类型是无缓存的，是同步阻塞的
//执行c<-true命令时，就是说，往channel  c里存true
//如果此时，没有读取的话，存true的时候，就是阻塞的
//也就是说，必须是我一存储true,那么就立马有人拿走true


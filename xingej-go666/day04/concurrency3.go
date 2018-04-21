//对Chanel的基本操作
package main

import "fmt"

func main() {

	c := make(chan bool) //创建一个bool类型

	// 其实，真正的业务逻辑，也就下面的方法，其他都是非业务逻辑，都是Go语言
	go func() {
		fmt.Printf("Go go go !")

		//将true 存储到c 里
		c <- true
	}()

	//主main程序，运行到这里后，会
	//进行阻塞的
	//直到c读取都值true
	<- c

}








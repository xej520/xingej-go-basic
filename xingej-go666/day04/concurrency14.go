//将select 用于发送数据  1
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	select {
	case v := <-c:
		fmt.Println(v)
	//3 秒后，下面的case 接收到chan类型的消息，就开始运行case
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")

	}
}

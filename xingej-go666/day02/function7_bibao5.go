package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i< 1000; i++{
		go func(i int) {
			//注意一下
			//println 方法，属于IO操作
			//进行IO操作时，总有切换等待的时间
			//进行切换时，会交出控制权，让其他go routines执行
			//如果始终不交出控制权的话，其他go routines就没法进行。
			fmt.Println("hello from goroutines, ", i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}


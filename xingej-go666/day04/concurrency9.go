package main

import (
	"fmt"
	"runtime"
)

func main() {
	//设置程序，按照多核程序运行
	//当按照多核处理时，goroutine 并不是按照顺序执行的，
	runtime.GOMAXPROCS(runtime.NumCPU())

	//创建一个无缓存的channel
	c := make(chan  bool)

	for i := 0; i<10 ;i++  {
		go sum(c, i)
	}

	//读取c的值，
	<- c
}

func  sum(c chan bool ,index int){
	a := 1
	for i := 0; i<10000000;i++   {
		a += i
	}

	fmt.Println(index, a)

	//在多核处理情况下，按照下标来判断是否所有的goroutine都运行结束，
	// 是不行的，因为很有可能首先运行的goroutine就是最后一个
	//如何解决这种现象呢？
	//两种方法：
	// 1、缓存解决
	//
	if  index == 9{
		c <- true
	}

}











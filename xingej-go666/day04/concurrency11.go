//waitgroup 来代替channel,
//打印出所有的goroutine
//waitgroup的工作原理？
//就是用来添加需要工作的任务，每完成一次任务就标记一次Done, 随之待完成量会减1，
//主线程根据任务组里的任务数量是否是0，来判断，如果是0的话，就说明所有的任务都全部完成了，
//主线程，可以结束了，从而实现了与阻塞队列类似的同步功能
package main

import (
	"sync"
	"fmt"
)

func main() {
	//创建一个空的waitGroup任务组
	wg := sync.WaitGroup{}

	//添加10个任务，说明有10个任务需要允许
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go getSparkInfo(&wg, i)
	}

	//主线程一直等待，直到检测到任务数为0
	wg.Wait()
}

// 定义的function getSparkInfo 在同一个包day04下是唯一的
//不能有相同的名字
//这里是引用拷贝，不是能是按值拷贝，如果是按值拷贝的话，不会修改旧值的
//这样就会发生死循环导致死锁
func getSparkInfo(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}

	fmt.Println(index, a)

	//表示，这个function已经完成了
	wg.Done()
}

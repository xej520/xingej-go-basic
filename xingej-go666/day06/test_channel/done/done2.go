//测试通过 通信来共享内存?
//目的，就是 当接收方完成后，通知发送方
package main

import "fmt"

func main() {
	chanDemo2()

}

func chanDemo2()  {
	var workers [10]worker2

	for i:=0; i<10; i++{
		workers[i] = createWorker2(i)
	}

	//开始给worker发送数据
	//先发送小写字母
	for i, worker := range workers{
		worker.receive <- 'a' + i
	}

	//开始接收
	for _, worker := range workers {
		<-worker.do
	}

	//开始发送大写字母
	for i, worker := range workers{
		worker.receive <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.do
	}

}

func doWorker2(index int, worker worker2)  {
	for value :=range worker.receive{
		//接收方，实现自己的业务逻辑
		fmt.Printf("Worker %d received %c\n", index, value)

		//接收方，完成后，告知发送方，已经完成了
		worker.do <- true
	}
}

//声明一类型
type worker2 struct {
	receive chan int
	do chan bool
}

func createWorker2 (index int) worker2 {
	worker := worker2{receive:make(chan int), do:make(chan bool)}
	go doWorker2(index, worker)

	return worker
}


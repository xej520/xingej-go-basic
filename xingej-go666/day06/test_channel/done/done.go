//测试通过 通信来共享内存?
//目的，就是 当接收方完成后，通知发送方
package main

import "fmt"

func main() {
	chanDemo()

}

func chanDemo()  {
	var workers [10]worker

	for i:=0; i<10; i++{
		workers[i] = createWorker(i)
	}

	//开始给worker发送数据
	for i, worker := range workers{
		worker.receive <- 'a' + i
	}

	for i, worker := range workers{
		worker.receive <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.do
		<-worker.do
	}

}

func doWorker(index int, worker worker)  {
	for value :=range worker.receive{
		//接收方，实现自己的业务逻辑
		fmt.Printf("Worker %d received %c\n", index, value)

		go func() {
			//接收方，完成后，告知发送方，已经完成了
			worker.do <- true
		}()
	}
}

//声明一类型
type worker struct {
	receive chan int
	do chan bool
}

func createWorker (index int) worker {
	worker := worker{receive:make(chan int), do:make(chan bool)}
	go doWorker(index, worker)

	return worker
}


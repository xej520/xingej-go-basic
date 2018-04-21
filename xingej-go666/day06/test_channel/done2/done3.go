package main

import (
	"sync"
	"fmt"
)

func main() {
	chanDemo3()
}

type appWorker struct {
	receive chan int

	do func()
}

func chanDemo3() {
	var wg sync.WaitGroup
	var workers [10]appWorker
	for i:=0; i<10; i++{
		workers[i] = createChanWorker(i, &wg)
	}
	wg.Add(20)
	for i,worker :=range workers{
		worker.receive <- 'a'+i
	}

	for i,worker :=range workers{
		worker.receive <- 'A'+i
	}

	wg.Wait()
}

func createChanWorker(id int, wg *sync.WaitGroup) appWorker {
	appW := appWorker{
		receive: make(chan int),
		//更加抽象了
		do: func() {
			wg.Done()
		},
	}
	go chanWorker(id, appW)

	return appW
}

func chanWorker(id int,  w appWorker) {
	for n := range w.receive {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.do()
	}
}

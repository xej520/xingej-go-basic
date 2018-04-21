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

	wg *sync.WaitGroup
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
		wg:      wg,
	}
	go chanWorker(id, appW.receive, appW.wg)

	return appW
}

func chanWorker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		wg.Done()
	}
}

package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

//测试wait.Poll的用法
func main() {
	wait.Poll(time.Second*1, time.Second*3, show1)
}

func show1() (bool, error) {
	i := 0
	for {
		fmt.Printf("第%d次打印 hello k8s!\n", i)
		time.Sleep(time.Second * 2)
		if i++; i > 5 {
			return true, nil
		}

	}
}

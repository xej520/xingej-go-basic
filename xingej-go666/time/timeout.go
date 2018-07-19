package main

import (
	"fmt"
	"time"
)

func main() {

	startTime := time.Now()

	// 休息5秒钟
	time.Sleep(time.Second * 5)

	endTime := time.Now()

	fmt.Printf("----经过了%d 时间啊", endTime.Second()-startTime.Second())

}

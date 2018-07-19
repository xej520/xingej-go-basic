package main

import (
	"fmt"
	"time"
)

func main() {
	old := time.Now()

	time.Sleep(time.Second * 5)

	new := time.Now()

	fmt.Println("----时间差:\t", int(new.Sub(old).Seconds()))
}

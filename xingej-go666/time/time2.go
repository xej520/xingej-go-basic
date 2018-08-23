package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	old := time.Now()

	time.Sleep(time.Second * 5)

	new := time.Now()

	fmt.Println("----时间差:\t", int(new.Sub(old).Seconds()))

	var clust map[string]string
	clust = make(map[string]string)
	clust["x.b"] = "laoxing"
	fmt.Println("====", clust["x.b"])

	fmt.Println("------>\t", strconv.FormatInt(int64(6633), 10))

}

//生成一个随机数
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("--->:\t", randInt(0, 1000))
}

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if min > max {
		return max
	}

	return r.Intn(max-min) + min
}

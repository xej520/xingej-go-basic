package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	var msg = "187 47 91 122 114 81 230 213 181 150 90 20 9 20 224 223"
	has := md5.Sum([]byte(msg))

	fmt.Println("---->\n", has)

}

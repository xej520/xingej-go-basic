package main

import "fmt"

func main() {
	sparkClusterIdApp := map[string]string{"spark001": "marathon001", "spark002": "marathon002"}

	//遍历map类型，
	//第一个参数，就是键
	//第二个参数，就是值
	for key, value := range sparkClusterIdApp {
		fmt.Println("key:\t", key, "\tvalue:\t", value)
	}

}

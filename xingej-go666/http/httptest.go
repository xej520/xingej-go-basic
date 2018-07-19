package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 模拟GET请求
	resp, err := http.Get("http://www.baidu.com")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}

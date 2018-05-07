package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlInfo, err := url.Parse("http://www.baidu.com")
	if err != nil {
		return
	}
	data := urlInfo.Query()

	fmt.Println(data)
}

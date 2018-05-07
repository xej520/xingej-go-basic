package main

import (
	"fmt"
	"regexp"
)

const text = "my email is k8sAndDocker@google.com"

func main() {
	//re 是 正则表达式的匹配器
	re, err := regexp.Compile("k8sAndDocker@google.com")
	if err != nil {
		panic(err)
	}
	result := re.FindString(text)
	fmt.Println("result:\t", result)
}

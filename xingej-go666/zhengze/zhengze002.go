package main

import (
	"fmt"
	"regexp"
)

const text_1 = "my email is k8sAndDocker@google.com"

func main() {
	//目前的正则表达式，仅仅是匹配一个值，k8sAndDocker@google.com
	re := regexp.MustCompile("k8sAndDocker@google.com")
	match := re.FindString(text_1)

	fmt.Println(match)
}

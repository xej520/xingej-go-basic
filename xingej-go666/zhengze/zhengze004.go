package main

import (
	"fmt"
	"regexp"
)

const text_4 = "my email is k8sAndDocker@google.com"

func main() {
	//只匹配小写字母，大写字母，数字，不允许有特殊符号
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text_4)

	fmt.Println(match)
}

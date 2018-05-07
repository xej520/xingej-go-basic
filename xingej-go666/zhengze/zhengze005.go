package main

import (
	"fmt"
	"regexp"
)

const text_5 = `
	my email is k8sAndDocker@google.com
	my email is spark@qq.com
	my email is hadoop@126.com
	my email is kafka@163.com
	my email is docker@163docker.com.cn
`

func main() {
	//在[]里,  . 不需要 添加 转义字符
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	//-1 表示，要匹配所有满足条件的词
	match := re.FindAllString(text_5, -1)

	fmt.Println(match)
}

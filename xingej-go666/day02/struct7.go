//匿名字段  测试
package main

import "fmt"

type students struct {
	//这些就是匿名字段，没有定义名字
	string
	int
}

func main() {
	boy := &students{
		//初始化的时候，必须按照顺序来进行的
		"xiaoqiang", 19}
	fmt.Println(boy)
}

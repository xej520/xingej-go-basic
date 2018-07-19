package main

import "fmt"

//就是测试append的一个特性
func main() {
	var names []string
	subnames := []string{"hello", " spark", " hadoop"}

	//添加一个切片，直接在后面，添加三个。。。就可以了
	names = append(names, subnames...)
	names = append(names, "---spark:", "192.168.1.110:7077")
	for _, value := range names {
		fmt.Println("===>:\t", value)
	}
}

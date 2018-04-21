//主要是想练习，type 的用处
package main

import "fmt"

//定义接口类型
type adder interface {
	add(string) int
}

//定义函数类型
type handler func(name string) int

//实现函数类型方法
func (h handler) add(name string) int{
	return h(name) + 10
}

//函数参数类型 接收 实现了adder接口的对象(函数或结构体)
func process( add adder) {
	fmt.Println("process:\t", add.add("k8s"))
}

//
func doubler(name string) int {
	return len(name) * 2
}

//非函数类型
type myInt int

//实现了adder接口
//（i myInt）是一个receiver，限制此方法，只能是哪个类来调用
func (i myInt) add(name string) int {
	return len(name) + int(i)
}

func main() {
	//注意，要成为函数对象，必须显示定义handler类型
	var myHandler handler = func(name string) int {
		return len(name)
	}

	//调用函数
	fmt.Println("--->调用函数:\t", myHandler("k8s"))
	fmt.Println("--->调用函数对象的方法:\t", myHandler.add("k8s"))
	fmt.Println("--->doubler函数显示转换成handler函数对象，然后再调用对象的add方法:\t", handler(doubler).add("k8s"))

	//以下是针对接口adder的调用
	process(myHandler) //process函数需要 adder接口类型参数
	process(handler(doubler)) //因为process接收的参数类型是handler，所以必须进行强制转换
	process(myInt(8))  //实现adder接口不仅可以是函数也可以是结构体

}














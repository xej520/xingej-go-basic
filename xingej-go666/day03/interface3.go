//不同接口之间的转换，类似于Java中的向上转型，或者向下转型
//就是说，有两个接口A， B
//其中，A接口里，嵌入了B
//那么A接口可以转换成B接口，但是，B接口不能转换成A接口，因为
//A接口里，可能包含B接口里没有的方法
//也就是说，多的可以向少的转换，反之不可。
package main

import "fmt"

type USB2 interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PcConnecter struct {
	name string
}

func (pcConnecter PcConnecter)Name() string {
	return pcConnecter.name
}

func(pcConnecter PcConnecter)Connect() {
	fmt.Println("Connected:\t", pcConnecter.name)
}

func main() {
	pc := PcConnecter{"appleConnector"}
	var a Connecter
	//将USE2 类型，强制 转换成了Connecter类型
	a = Connecter(pc)
	a.Connect()

	//------验证----只有当接口存储的类型和对象都为nil时，接口才等于nil
	//声明一个空接口, 也没有赋值，就是nil
	var aa interface{}
	fmt.Println( aa == nil)  // aa 本身就是一个nil,啥也没存

	//变量p， 是一个指向int类型的指针
	//直接初始化为nil
	var p *int = nil
	aa = p //aa 指向p
	fmt.Println( aa == nil)
}


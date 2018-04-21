// interface 接口  练习
//实现接口的原则就是
//实现了它定义的方法，就默认是实现了接口
package main

import "fmt"

//声明一个接口
type USB interface {
	//声明方法Name,  并设置 返回值类型string
	Name() string

	//声明方法Connect方法，无返回值
	Connect()
}

//声明一个类型，在Go语言中，对应的就是struct类型
type PhoneConnector struct {
	//声明一个私有属性
	name string
}
//---------------------声明完Name,Connector方法后，就是实现了USB接口了---------------------------
//使用receiver，将类型跟方法进行绑定
func (pc PhoneConnector)Name() string{

	return pc.name
}

func (pc PhoneConnector)Connector()  {
	fmt.Println("connected:\t", pc.name)
}

//--------------------------------------------------------------------------------------
func main() {
	a := PhoneConnector{"apple"}
	fmt.Println("Name:\t", a.Name())
	a.Connector()
}







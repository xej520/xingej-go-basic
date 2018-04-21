//interface 嵌套  练习测试
package main

import "fmt"


//定义一个空的接口
//这样的话，所有的类，都默认实现了这个接口，因为它没有方法
//定义的空接口，就相当于定义了一 个Object对象，最高层
//都是它的子类了，就没有任何限制了
type nullEmpty interface {

}


// 定义一个  父接口
type HOME interface {
	//这个接口里，只定义一个方法
	Name() string
}

//再定义一个接口, 这相当于是子接口了
type MyHome interface {
	Show() string
	//这样就嵌套了 一个接口
	HOME
}

//声明一个结构类型
type BeijingHome struct {
	name string
}

//-----------开始创建方法method-----
func (info BeijingHome)Show(){
	fmt.Println("Show()--->info.name:\t", info.name)
}

func (info BeijingHome)Name() string{

	return info.name
}

func main() {
	a := BeijingHome{"yihuyuan"}

	fmt.Println("name:\t", a.Name())
	a.Show()

	//------下面演示一下，上面理论中说的---复制品的问题----
	b := a
	b.name = "lenovo" //修改后，并没有修改a里的值
	b.Show()

	Disconnect(a)
	Disconnect2(a)
}

//设计一个简单的类型断言
func Disconnect(home HOME){
	if pc, ok := home.(BeijingHome); ok{
		fmt.Println("Disconnected:\t", pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}

//设计一个简单的类型断言
//传入的参数，是，空接口
//实际上，对于传入的参数，就没有限制了
func Disconnect2(home interface{}){
	//对Disconnect()方法，进行改造，
	//因为，传入的参数是顶层，相当于Java里的Object，没有任何限制
	//类型，需要自己判断
	switch v := home.(type) {
	case BeijingHome:
		fmt.Println("Disconnected:\t", v.name)
	default:
		fmt.Println("Unknown decive.")
	}

}


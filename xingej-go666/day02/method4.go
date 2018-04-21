//方法 的权限问题，是否可以访问 私有字段呢？
//权限是以package 为级别的
//方法的访问权限是很高的，可以访问同一个package下的所有属性，包括公共的，私有的
package main

import "fmt"

type bike struct {
	//小写是私有属性，私有属性，只能在同一个package内，进行访问的
	name string
	//大写是公共属性
	Color string
}

func main() {
	bike := bike{name:"捷安特", Color:"黑色"}
	fmt.Println("old bike:\t", bike)

	//通过类型的变量，来调用方法
	bike.show()
	fmt.Println("new bike:\t", bike)
}

func (bike *bike)show()  {
	bike.Color = "红色"
	bike.name = "永久" //看见了吧，方法是可以访问同一package下的私有属性的
	fmt.Println("bike:\t", bike)
}













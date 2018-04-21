//创建方法
package main

import "fmt"

type book struct {
	Name string
}

type eleBook struct {
	Name string
}

func main() {
	//创建一个对象
	aBook := book{Name:"spark computer"}
	aBook.show()

	bBook := eleBook{Name:"k8s computer"}
	bBook.show()
}

//创建一个method
//此mothod方法属于book类型
//(a book)  这个字段 是说
// 此方法，属于哪类型，是跟类型绑定的
//只能指定的类型的变量，才能调用
//方法的绑定，只能是同一个包中，才起作用
func (a book) show()  {
	fmt.Println("this is book:\t",a.Name)
}

// 这属于重载了
//func (a book) show(info string)  {
//	fmt.Println("this is book")
//}

//创建一个method
//此mothod方法属于eleBook类型
func (a eleBook) show()  {
	fmt.Println("this is ele book")
}


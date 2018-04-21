//
package main

import "fmt"

//定义了一 个空的结构体
type st01 struct {
}

//定义一个非空的结构体
type person struct {
	Name string
	Age  int
}

func main() {
	personInfo := person{}
	personInfo.Age = 19
	personInfo.Name = "xiaoqiang"

	fmt.Println(personInfo)

	//------------------------------
	fmt.Println("=============================")
	//直接初始化   属于 字面值 初始化
	personInfo2 := person{Name :"spark01", Age: 19}
	fmt.Println(personInfo2)

}

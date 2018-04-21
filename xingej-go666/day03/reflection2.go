//反射 是如何处理 匿名字段的？
package main

import (
	"reflect"
	"fmt"
)

type Stu struct {
	Id int
	Name string
	Age int
}

type Man struct {
	//这里你要注意一下，你创建的属性，是有顺序的，是有下标的
	//如Stu 下标 就是0， title下标就是1
	// Stu 就是匿名属性
	Stu
	title string
}

func main() {
	//注意，对匿名字段进行初始化时的方式，其实本质上跟其他属性是一样的
	m := Man{Stu:Stu{Id:2,Name:"Jack",Age:19}, title:"Manager"}
	t := reflect.TypeOf(m)

	//取匿名字段的方式
	//FieldByIndex 方法，传入的是一个切片slice类型
	//第1个0，表示，匿名字段在Man中的下标位置
	//第2个0，表示，你要取匿名字段中哪个属性的下标
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,0})) //取的是id
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,1})) //取的是Name
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,2})) //取的是Age

}








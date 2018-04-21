//结构体，实现 类似于 继承的效果
package main

import "fmt"

type anminal struct {
	//设置一些共有的属性
	Name, address string
}

type cat struct {
	//anminal  Go 语言，默认，anminal是类型，同时也是属性名称
	anminal
	Sex int			// 猫的特有属性，性别是啥
}

type dog struct {
	anminal
	Hobby string  //狗的特有属性，爱好
}

func main() {
	//第一种初始化方式
	xiaoCat := cat{Sex:0, anminal : anminal{Name:"xiaohong", address:"beijing"}}
	xiaoDog := dog{Hobby:"play", anminal:anminal{Name:"xiaohuang", address:"shanghai"}}

	//第二种初始化方式
	xiaoCat.anminal.Name = "xiaoxiaoxiaohong" //这种方式，是为了防止名字相同时，冲突
	xiaoCat.Name = "xiaoxiaohong"

	fmt.Println("cat:\t", xiaoCat)
	fmt.Println("dog:\t",xiaoDog)

}





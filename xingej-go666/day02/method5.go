//练习题：
//依据的理论知识：为结构增加方法的知识
//尝试声明一个底层类型为int的类型
//并实现调用某个方法就递增100
//如   a:=0, 调用a.Increase()只会，a从0变成100
package main

import "fmt"

//为int 声明一个别名
type intFor int

func main() {
	var a intFor = 0
	fmt.Println(a.Increase())          //100
	fmt.Println(a.IncreaseV2(120))  //220
}

//传的地址
// 注意，返回值的类型是intFor
func (arg *intFor)Increase()  intFor {
	//添加上*,就表示取地址里的内容了
	*arg = 100

	return *arg
}

func (arg *intFor)IncreaseV2(num int) intFor {

	*arg += intFor(num)

	return *arg
}


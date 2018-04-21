//别名 与 方法的绑定
// 作用，就是，有一种增强的感觉，如int类型，本身没有show,add方法
//别名  与   方法的组合，确有了
package main

import "fmt"

type TZ int

func main() {
	var ty TZ = 3
	fmt.Println("ty:\t" , ty)
	//Method value 调用方式，通过类型的变量来调用
	ty.show()
	ty.add()

	fmt.Println("==================Method value   Method Express的不同==============================")
	//对方法的两种不同的导调用方式而已

	//Method express 调用方式，通过类型直接来调用
	//此种方式，需要自己输入 变量
	//因为receiver接收的是地址，因此，我们传入的是地址
	(*TZ).show(&ty)
	(*TZ).add(&ty)


}

//下面的这些方法，都是通过 某一个类型的 变量 来进行调用的
//java 是通过对象来调用的
func (tz *TZ)show()  {
	fmt.Println("---->:\t 这是int类型", *tz) //tz是地址，*tz 是取地址里的内容了
}

func (tz *TZ)add()  {
	fmt.Println("---->:\t 这是int类型的加法")
}








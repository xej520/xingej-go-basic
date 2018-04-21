// goto语句 练习，需要配合标签来使用
// goto 是调整 程序的执行位置，而不是像continue，继续执行
// 因此，使用goto 的时候，最好将标签，放在for循环的下面，尽量避免死循环
package main

import "fmt"

func main() {
	//方式一：
	//跳到标签LABEL1的位置，继续执行，并没有结束循环
	//定义一个无限循环
	for {
		//使用一个有限循环
		for i := 1; i < 10; i++ {
			if i > 5 {
				fmt.Println("i > 5")
				goto LABEL2
			}
		}

	}
	// 也可以将标签放在否循环的后面
LABEL2:
	fmt.Println("跳出循环了么？")
fmt.Println("--------------------不建议使用下面的方式，有可能引起-----------死循环--------------------")
	//方式二：
	//跳到标签LABEL1的位置，继续执行，并没有结束循环
LABEL1:
//定义一个无限循环
	for {
		//使用一个有限循环
		for i := 1; i < 10; i++ {
			if i > 5 {
				fmt.Println("i > 5")
				goto LABEL1
			}
		}

	}
	fmt.Println("跳出循环了么？")

}

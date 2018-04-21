//continue 结合标签 练习
package main

import "fmt"

func main() {
	LABEL1:
		for i:=1; i< 10;i++  {
			for{
				fmt.Println("i:\t", i)
				//跳出当前的无限循环，继续执行外层的循环
				goto LABEL1
			}
		}

	fmt.Println("=====> 结束循环！")
}




//break 结合标签的使用
package main

import "fmt"

func main() {
	LABEL1:
		for {
			for i:=1;i<10 ;i++  {
				if i>5 {
					fmt.Println("i > 5")
					//break ,跳出跟标签LABEL1同层的循环
					break LABEL1
				}
			}
		}

	fmt.Println("=====>:\t 结束循环了没？" )

	for i:=2; i< 4;i++ {
		LABLE2:
			for {
				if i >= 2 {
					//break会跳出跟LABEL2同层的循环
					break LABLE2
				}
			}

		fmt.Println("i:\t", i)
	}


}















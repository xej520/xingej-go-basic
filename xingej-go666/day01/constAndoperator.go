//此文件用来练习
//目的:结合常量iota与<< 运算符实现计算机存储单位的枚举
// byte, KB,MB,GB,TB,PB等等
package main

import "fmt"

const(
	byte float64 = 1 << (iota * 10)
	//说明：
	//下面的常量，并没有初始化，就会默认采用上面的表达式
	// KB = 1 << (iota * 10)， 只是，iota 再不断的增加
	KB
	MB
	GB
)

func main() {
	fmt.Println(byte)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}

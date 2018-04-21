//综合练习
package main

import (
	"os"
	"fmt"
	"bufio"

	"xingej-go666/day00/defer_test/fib"
)

func main() {
	writeFile("fib.txt")
}

func writeFile(filename string) {
	file, error := os.Create(filename)
	//校验，文件是否创建成功
	if error != nil {
		fmt.Errorf("创建文件%s失败", filename)

		//执行panic后，就结束程序运行
		//有点类似于 java 中 throw，抛出异常
		panic(error)
	}

	//声明一个defer，用于关闭资源，如关闭文件
	defer file.Close()

	//将写到缓存里，提高效率
	writer := bufio.NewWriter(file)
	//从内存，刷新到磁盘文件里
	defer writer.Flush()

	f := fib.Fibonacci()
	for i:=0; i<20; i++{
		fmt.Fprintln(writer, f())
	}

	//总之，先运行的是writer.Flush()， 然后，再去运行file.Close()，defer的作用就发挥出来了，
	//是不是很合理
}

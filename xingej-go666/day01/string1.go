//struct 构建一个树 练习测试
package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func main() {
	var root treeNode
	root = treeNode{value:3}
	//仔细看，是不是有点类似于不同的构造函数
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.left.right = new(treeNode)
	root.right.left = createNode(3)


	fmt.Println("treeNode:\t", root)

	//定义一个slice
	nodes := []treeNode{
		{value:3},{5,nil,nil}	}
	fmt.Println("nodes:\t", nodes)

}

//创建一个工厂函数，
func createNode(value int) *treeNode{
	//返回的是个局部变量
	//那么，既然是局部变量? 其他地方能引用么？ 局部变量，不是说，函数结束后就结束了么
	//这是Go 语言的特殊，如果返回值的地址其他地方引用的话，就将这个变量，创建在堆上
	//如果没有引用的话，就将这个变量，创建在栈上。参与垃圾回收
	return &treeNode{value:value}
}


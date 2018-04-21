// 数组 练习
package main

import "fmt"

func main() {

	//定义数组时，必须指定 数组的个数
	var age [2]int
	fmt.Println(age[0])

	// 定义数组时，并且初始化元素
	sparkClusterName := [3]string{"spark-bj", "spark-shanghai", "spark-shenzhen"}
	fmt.Println(sparkClusterName[0])
	fmt.Println(sparkClusterName[1])
	fmt.Println(sparkClusterName[2])
	fmt.Println(sparkClusterName) // 直接将全部内容，以列表的形式，打印出来, [spark-bj spark-shanghai spark-shenzhen]

	// 初始化时，可以给指定位置元素，设定值，
	// 其他位置，没有设置的话，会自动添加默认值
	//如，下面的例子，仅仅对下标为1,4,9的位置，进行了设置，其他位置还是0
	sparkClusterCpu := [10]float64{1: 3.4, 4: 5.6, 9: 8.8}
	fmt.Println(sparkClusterCpu) //[0 3.4 0 0 5.6 0 0 0 0 8.8]

	// go语言 可以自动推算出数组的个数
	// 必须添加三个点 ...， 来表示数组个数
	// 方式一:已经知道 每个元素的值
	appleList := [...]int{1, 3, 4, 5}
	fmt.Println(appleList) //[1 3 4 5]
	//方式二:利用下标索引，来实现初始化
	orangeList := [...]int{0: 1, 1: 3, 2: 5}
	fmt.Println(orangeList) //[1 3 5]
	//方式三：利用最大下标，来实现数组个数的推算，go语言根据最大下标，来创建数组个数
	// 如，只设置下标为9的值为8，其他还是默认值
	bananaList := [...]int{9: 8}
	fmt.Println(bananaList) //[0 0 0 0 0 0 0 0 0 8]
}

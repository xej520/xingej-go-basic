//if控制语句 练习测试
package main

import "fmt"

func main() {
	age := 13
	if age < 18 {
		fmt.Println("未成年了")
	}

	fmt.Println("---------------------------------------------")
	//支持初始化表达形式, 就是先初始化，再写条件
	// 注意，money 是局部变量，仅仅作用域if代码块内
	if money := 2000; money < 5000 {
		fmt.Println("贫困户")
	}
	fmt.Println("---------------------------------------------")

	if zk_node_num, spark_node_num := 3,5; zk_node_num + spark_node_num > 10 {
		fmt.Println("zk节点数量与spark节点数量之和大于10!")
	} else {
		fmt.Println("zk节点数量与spark节点数量之和小于10!")
	}

	fmt.Println("---------------------------------------------")

	//变量以及 作用域说明
	var k8s_num = 15
	//当if内的局部变量与外部变量相同时，不会影响if外部的变量的值
	if k8s_num := 10 ; k8s_num > 8 {
		k8s_num := 6
		fmt.Println("if语句块内的k8s_num:\t" , k8s_num)  //6
	}

	//没有受到if中变量的影响
	fmt.Println("if语句块内的k8s_num:\t" , k8s_num) //15

}

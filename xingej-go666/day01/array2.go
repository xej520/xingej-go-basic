//数组的指针   指针的数组   练习
package main

import "fmt"

func main() {
	//初始化一个数组
	sparkClusterNameList := [...]string{"spark-beijing", "spark-shanghai", "spark-sjz"}
	//数组的指针
	var sparkP *[3]string = &sparkClusterNameList
	fmt.Println(sparkP[0]) //spark-beijing
	fmt.Println(sparkP[1]) //spark-shanghai
	fmt.Println(sparkP[2]) //spark-sjz
	fmt.Println(sparkP)    //&[spark-beijing spark-shanghai spark-sjz]

	//--------------------------数组的指针--------------------------------------
	// 同时声明多个变量
	age, num := 19, 1
	person := [...]*int{&age, &num}
	fmt.Println(person)

	// 同时，创建两个变量
	//name, school := "xiaoming","qinghua"
	//// 初始化一个数组info
	//info := [...]*int{&name, &shcool}
	//fmt.Println(info)

	//--------------------------数组的比较-------------------
	// 数组之间，只能进行  是否相等==运算，或者  不等!= 运算
	// 不能进行  大于> ，小于< 比较
	// 注意，数组进行比较，必须在个数相同的情况下，也就是说类型相同的情况下
	// 在Go语言中，如果数组的个数不相同的，就是不同的类型
	//创建3个数组
	appList := [...]int{3, 4, 5}
	bananaList := [...]int{3, 4, 5}
	orangeList := [...]int{3, 4, 6}
	fmt.Println(appList == bananaList) //true
	fmt.Println(appList == orangeList) //false

	// =======数组---以及----指向数组的指针---------
	// 都可以单独通过下标，进行设置单个元素
	kafkaArray := [10]int{9: 2}
	kafkaArray[2] = 8
	// 可以通过new关键字，来进行创建 数组的指针
	zkArray := new([10]int)
	zkArray[2] = 9
	fmt.Println(kafkaArray)
	fmt.Println(zkArray)

	//===============多维数组======================
	sparkInfo := [2][3]int{
		{2, 3, 1},
		{3, 4, 2}}

	fmt.Println(sparkInfo)

	//=============Go语言版本的====冒泡排序========
	studentScore := [...]int{5, 4, 2, 9, 1, 7}
	fmt.Println("==========排序前==========\n", studentScore)

	len := len(studentScore)

	for i := 1; i < len; i++ {
		for j := len - i; j > 0; j-- {
			if studentScore[j] > studentScore[j-1] {
				temp := studentScore[j-1]
				studentScore[j-1] = studentScore[j]
				studentScore[j] = temp
			}
		}
	}

	fmt.Println("==========排序后==========\n", studentScore)

}

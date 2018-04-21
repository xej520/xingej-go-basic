//结构体，按值 传递
package main

import "fmt"

//定义一个类型
type sparkInfo struct {
	SparkClusterName string
	SparkNodeNum     int
}

func main() {
	newSparkInfo := sparkInfo{
		SparkNodeNum:     8,
		SparkClusterName: "spark-test-001"}

	fmt.Println("main:\t", newSparkInfo)
	updateSparkInfo(newSparkInfo)
	fmt.Println("main:\t", newSparkInfo)

}

//同样，这里是值传递，内部的修改，并不会影响到 旧值的
func updateSparkInfo(sparkInfo sparkInfo) {
	sparkInfo.SparkNodeNum = 9
	fmt.Println("updateSparkInfo:\t", sparkInfo)
}

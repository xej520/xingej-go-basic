//将匿名结构，嵌套进别的结构体里
package main

import "fmt"

type hadoop struct {
	HadoopClusterName string
	HadoopClusterNum  int
	//创建一个匿名结构
	HadoopOtherInfo struct {
		//同样，当多个变量都一样的时候，也可用省略
		//这是Go语言的优点
		HadoopVersion, HadoopUrl string
	}
}

func main() {
	hdfs := &hadoop{
		HadoopClusterName: "Hadoop-test-001",
		HadoopClusterNum:  9}

	//只能通过这种方式，进行初始化
	hdfs.HadoopOtherInfo.HadoopUrl = "http://192.168.1.110:50070"
	hdfs.HadoopOtherInfo.HadoopVersion = "v2.7.0"

	fmt.Println(hdfs)
	fmt.Println(*hdfs)
	fmt.Println("HadoopClusterName:\t", hdfs.HadoopClusterName)
	fmt.Println("HadoopClusterNum:\t", hdfs.HadoopClusterNum)
	fmt.Println("HadoopClusterVersion:\t", hdfs.HadoopOtherInfo.HadoopVersion)
	fmt.Println("HadoopClusterUrl:\t", hdfs.HadoopOtherInfo.HadoopUrl)

}

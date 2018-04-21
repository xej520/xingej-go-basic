//测试某个值是否在map里呢？
package main

import "fmt"

func main() {

	sparkClusterIdApp := map[string]string{"spark001":"marathon001","spark002":"marathon002"}

	clusterId001 := "spark003"

	_, ok := sparkClusterIdApp[clusterId001]

	if !ok {
		fmt.Println("clusterId: \t", clusterId001, "不再容器里")
	}

	clusterId002 := "spark001"

	if v,ok :=sparkClusterIdApp[clusterId002]; ok {
		fmt.Println("value:\t", v)
	} else {
		fmt.Println("value does not exist!")
	}

}



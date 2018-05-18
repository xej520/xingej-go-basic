package main

import "fmt"

type SparkPod struct {
	Name string
	//声明一个匿名方法
	// 声明匿名方法的好处，就是
	//不同的实例，有不同的实现方法
	Show func(key string) error
}

func main() {
	sparkPod := SparkPod{
		Name: "spark-k8s-cluster",
		Show: func(key string) error {
			if len(key) > 10 {
				fmt.Println("------>:\t", len(key))
				return nil
			} else {
				return nil
			}
		},
	}

	sparkPod.Show("hello, spark, k8s, docker")
	//其实，就是获取的是属性
	hel := sparkPod.Show

	//然后，通过属性，来调用方法
	hel("dafddfafdafd")
}

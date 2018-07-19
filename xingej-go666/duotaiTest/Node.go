package main

//模拟java中的接口，实现类

import "fmt"

// 声明一个接口，
type workthread interface {
	do(name string)
}

// 传入的wt的子类不同，do方法执行的逻辑也不同
func Handle(wt workthread, name string) {
	fmt.Println("--------执行do方法前的逻辑------")
	wt.do(name)
	fmt.Println("--------执行do方法后的逻辑------")
}

type SparkNode struct {
	ClusterName string
}

func (sn *SparkNode) do(name string) {
	fmt.Println("----hello, I'm spark node", "\t", sn.ClusterName)
}

func (sn *KafkaNode) do(name string) {
	fmt.Println("----hello, I'm kafka node\t", sn.ClusterName)
}

type KafkaNode struct {
	ClusterName string
}

func main() {

	sparkNode := &SparkNode{
		ClusterName: "k8s",
	}

	kafkaNode := &KafkaNode{
		ClusterName: "kafka-node",
	}

	Handle(sparkNode, "111")
	Handle(kafkaNode, "222")

}

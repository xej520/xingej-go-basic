//map 练习
package main

import (
	"fmt"
	"sort"
)

func main() {
	//============================方式一：===========================
	//创建一个空的map
	//先声明一个map类型
	var nodeMap map[int]string
	//初始化nodeMap, 全部为空
	nodeMap = map[int]string{}
	fmt.Println(nodeMap)

	fmt.Println("-----------------------------------------------")
	//============================方式二：===========利用make================
	var clusterMap map[int]string = make(map[int]string)
	fmt.Println(clusterMap)

	fmt.Println("-----------------------------------------------")
	//============================方式三：===========利用推断功能================
	operatorMap := make(map[int]string)
	fmt.Println(operatorMap)

	//============================方式四：===========创建时直接初始化================
	//此种方式，不需要使用make
	operatorMap2 := map[int]string{3: "hello", 5: "world"}
	fmt.Println("===>:\t", operatorMap2)

	// ======初始化=====
	nodeMap[1] = "sparkNode"
	nodeMap[2] = "esNode"
	fmt.Println(nodeMap)

	//=======根据键 取出 元素
	nodeName := nodeMap[2]
	fmt.Println(nodeName)

	//======删除键值对
	delete(nodeMap, 1) //根据键值对，进行删除
	fmt.Println(nodeMap)

	fmt.Println("-----------------------------------------------")
	//--------------复杂map的操作-------------------
	//声明一个map类型
	var clusterAppTaskId map[string]map[string]string
	//初始化此map类型
	clusterAppTaskId = make(map[string]map[string]string)

	taskId, ok := clusterAppTaskId["spark-beijing"]["/spark-beijing/app-uewqr"]

	if !ok {
		//每一级别的map都有进行初始化，编译时是找不到的，只有运行时，才能发现
		clusterAppTaskId["spark-beijing"] = make(map[string]string)
	}

	clusterAppTaskId["spark-beijing"]["/spark-beijing/app-uewqr"] = "app-ewr-spark-taskid-001"
	taskId, ok = clusterAppTaskId["spark-beijing"]["/spark-beijing/app-uewqr"]

	fmt.Println(taskId, ok)

	fmt.Println("-----------------------------------------------")
	//--------------迭代操作-------------------
	//for2 i, v := range slice {
	//
	//}
	//i, 表示下标，v表示对应的值，是拷贝的值
	//要特别注意，对v的任何修改，都不影响原值，
	// map类型也是，不会影响原值的

	//例如，下面的例子，就是对v的操作后，不会对sm产生影响的
	//因此，不建议使用这种方式
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string)
		v[1] = "ok"
		fmt.Println(v)
	}

	fmt.Println(sm)
	fmt.Println("-----------------------------------------------")
	//下面的修改，原值
	for i := range sm {
		sm[i] = make(map[int]string)
		sm[i][2] = "spark"
		fmt.Println(sm[i])
	}

	fmt.Println(sm)

	fmt.Println("-----------------------------------------------")
	//---------------------------------
	//map是无序的，如何按照键从小到大获取map中的值
	//需要生成一个切片，来存储map的键
	//将键按照从小到大排序，然后，再根据键去取值
	marathonApp := map[int]string{1: "spark", 3: "es", 8: "ftp", 7: "hadoop", 4: "k8s", 2: "docker"}
	len := len(marathonApp)
	//生成一个切片，来存储键
	kSlice := make([]int, len)
	// 声明一个计数器，用于初始化切片时使用
	var i int = 0
	for k, _ := range marathonApp {
		kSlice[i] = k
		i++
	}
	fmt.Println("键排序前:\t", kSlice)

	//切片是引用传递，因此，下面排序后，不需要返回值进行接收
	sort.Ints(kSlice)
	fmt.Println("键排序后:\t", kSlice)

	fmt.Println("根据键按照从小到大，依次取出对应的值")
	//下面，开始迭代marathonApp，就可以按照键的从小到大，依次取出值了
	for _, v := range kSlice {
		//这里一定要注意，是用值，而不是 kSlice的下标
		fmt.Println(marathonApp[v])
	}

	fmt.Println("-----------------------------------------------")
	//------------------如何将map类型的键与值对换呢---------------
	kafkaClusterIdNameMap := map[int]string{1: "kafka-beijing", 2: "kafka-sjz", 4: "kafka-shanghai", 5: "kafka-tianjin", 3: "kafka-yizhuang"}
	kafkaNameClusterIdMap := make(map[string]int)
	for k, v := range kafkaClusterIdNameMap {
		kafkaNameClusterIdMap[v] = k
	}
	fmt.Println("kafkaClusterIdNameMap:\t", kafkaClusterIdNameMap)
	fmt.Println("kafkaNameClusterIdMap:\t", kafkaNameClusterIdMap)

}

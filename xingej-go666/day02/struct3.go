//结构体，按地址传递
package main

import "fmt"

type k8sInfo struct {
	K8sClusterName string //k8s集群的名称
	K8sClusterNumm int    //k8s集群的节点个数
}

func main() {
	k8sInfo := k8sInfo{
		K8sClusterNumm: 1,
		K8sClusterName: "k8s-test-001"}

	fmt.Println(k8sInfo)
	updateK8sClusterInfo(&k8sInfo)
	fmt.Println(k8sInfo)
}

//传递的是 地址，按地址传递
//修改了，旧的值
func updateK8sClusterInfo(info *k8sInfo) {
	info.K8sClusterNumm = 110

	fmt.Println("updateK8s:\t", info)
}

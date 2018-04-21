//结构体，按地址传递
package main

import "fmt"

type dockerInfo struct {
	DockerClusterName string
	DockerClusterNum int
}

func main() {
	//一般更习惯这种写法，
	//在初始化的时候，就直接获取地址
	//这样以后在调用的时候，就不需要添加&号了
	dInfo := &dockerInfo{
		DockerClusterNum:19,
		DockerClusterName:"docker-yizhuang"}

	fmt.Println("init docker:\t", dInfo)
	updateDockerInfo(dInfo)
	fmt.Println("after docker:\t", *dInfo)


}

func updateDockerInfo (info *dockerInfo) {
	info.DockerClusterNum = 80
	fmt.Println("udpateDocker:\t", info)
}




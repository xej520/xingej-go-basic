//匿名结构
package main

import "fmt"

func main() {

	//创建一个匿名结构，
	//并且进行了初始化
	//而且，直接获取地址&
	ftp := &struct {
		FtpName string
		FtpNum int
	}{
		FtpName:"ftp-beijing",
		FtpNum:8}

	fmt.Println(ftp)
	fmt.Println(*ftp)
	fmt.Println("FtpName:\t", ftp.FtpName)
	fmt.Println("FtpNum:\t", ftp.FtpNum)

}




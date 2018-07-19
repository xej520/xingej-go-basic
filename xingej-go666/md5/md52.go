package main

import (
	"crypto/md5"
	"fmt"
)

//注意，md5是不逆的
//md5的格式，基本上是一样的
//使用时，按照下面的格式，copy就可以了。
func main() {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte("admin"))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}

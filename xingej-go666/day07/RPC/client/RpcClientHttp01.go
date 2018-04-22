package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	//此时，定义的类型开头必须是大写的
	// 底层进行反射时，才能获取到，如果是小写的话，反射不到的。
	A, B int
}

type Quotient struct {
	//如8/5 的话，Quo就是1， Rem就是3
	Quo int //表示，整数
	Rem int //表示，余数
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:\t", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddr := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddr+":1234")

	if err != nil {
		log.Fatal("dialing, ", err)
	}

	log.Println("-----成功连接上----")

	args := Args{3, 4}
	var reply int
	err = client.Call("Math.Mul", args, &reply)

	if err != nil {
		log.Fatal("Math error:", err)
	}

	fmt.Printf("Math: %d * %d = %d\n", args.A, args.B, reply)

	fmt.Println("===============================")

	var quo Quotient

	err = client.Call("Math.Divide", args, &quo)
	if err != nil {
		log.Fatal("Math error:", err)
	}

	fmt.Printf("Math: %d/%d=%d remainder %d\n", args.A, args.B, quo.Quo, quo.Rem)
}

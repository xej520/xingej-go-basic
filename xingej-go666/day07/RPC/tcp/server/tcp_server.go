package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

func main() {
	math := new(Math)

	rpc.Register(math)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")

	if err != nil {

		fmt.Println("Fatal error:\t", err)
		fmt.Println(err.Error())
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("Fatal error: \t", err)
		os.Exit(2)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn error:\t", err)
		}
		rpc.ServeConn(conn)
	}

}

type Args struct {
	//此时，定义的类型开头必须是大写的
	// 底层进行反射时，才能获取到，如果是小写的话，反射不到的。
	A, B int
}

type Math int

func (m *Math) Mul(args *Args, reply *int) error {
	*reply = args.A * args.B

	return nil
}

type Quotient struct {
	//如8/5 的话，Quo就是1， Rem就是3
	Quo int //表示，整数
	Rem int //表示，余数
}

func (m *Math) Divide(args *Args, quo *Quotient) error {

	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B //求余数

	return nil
}

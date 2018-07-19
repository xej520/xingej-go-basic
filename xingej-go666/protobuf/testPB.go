package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"xingej-go/xingej-go/xingej-go666/protobuf/test"
)

func write() {
	p1 := &test.Person{
		Id:   1,
		Name: "spark",
	}

	book := &test.ContactBook{}
	book.Persons = append(book.Persons, p1)

	//编码数据
	data, _ := proto.Marshal(book)

	//将数据写入文件
	ioutil.WriteFile("./pb_test.txt", data, os.ModePerm)
}

func read() {
	data, _ := ioutil.ReadFile("./pb_test.txt")
	book := &test.ContactBook{}
	//解码数据
	proto.Unmarshal(data, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
	}
}

func main() {
	write()

	read()

}

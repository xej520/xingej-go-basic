package main

import (
	"encoding/json" //自带的json工具
	"fmt"
)

//github.com/pquerna/ffjson/ffjson
// 这个json工具，比自带的json效率要高很多
// 而且接口，完全一样，
type worker struct {
	Name  string `json:"worker_name"` //转换成json格式时，key就是worker_name
	Sex   string
	salay int32 //注意，json时，不会对这个属性操作，因为salay是小写
}

func main() {
	// 1、对数组形式，进行编码
	x := [5]int{1, 2, 3, 4, 5}
	bytes, e := json.Marshal(x)

	if e != nil {
		panic(e)
	}
	fmt.Println(string(bytes)) //[1,2,3,4,5]

	//	2、对map类型，进行json
	var stu = map[string]string{}
	stu = make(map[string]string)
	stu["name"] = "beijing"
	stuJson, err := json.Marshal(stu)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(stuJson)) //{"name":"beijing"}
	// 3、对struct对象进行JSON格式 编码
	wk := worker{
		Name:  "xiaozhang",
		Sex:   "male",
		salay: 34,
	}
	wrJson, err2 := json.Marshal(wk)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(string(wrJson)) //{"worker_name":"xiaozhang","Sex":"male"}

	//	4、对wrJson进行解码
	var w interface{} //声明一个interface， 用于存储解码后的值
	json.Unmarshal(wrJson, &w)
	fmt.Printf("---json格式解码---->%v", w) //map[worker_name:xiaozhang Sex:male]

}

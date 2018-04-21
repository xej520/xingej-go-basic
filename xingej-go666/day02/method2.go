package main

import "fmt"

type flower struct {
	Name string
}

type rose struct {
	Name string
}

func main() {

	flower := flower{Name:"a big flower"}
	rose := rose{Name:"a big rose"}
	fmt.Println("init flower:\t", flower)
	flower.show()
	fmt.Println("after flower:\t", flower)

	fmt.Println("===========================")

	fmt.Println("init rose:\t", rose)
	rose.show()
	fmt.Println("after rose:\t", rose)
}

//(flower flower)  这种方式，是按值传递的，不能改变原值的
func (flower flower)show()  {
	flower.Name = "I'm flower"
	fmt.Println("flower:\t", flower)
}

//(rose *rose) 是按引用传递的，可以改变原值的
func (rose *rose)show()  {
	rose.Name = " this is rose"
	fmt.Println("rose:\t", rose)
}





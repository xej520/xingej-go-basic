//通过反射，进行方法的调用，相当于动态调用了
package main

import (
	"fmt"
	"reflect"
)

type Teacher struct {
	Id int
	Name string
	Age int
}

//通过receiver将Show方法，跟Teacher类型，进行绑定
func (teacher Teacher)Show(name string) {
	fmt.Println("hello, ", name, ", my name is ", teacher.Name)
}

//注意======目前没有发现====如何通过====反射===来获取=====私有方法
func (teacher Teacher)info(){
	fmt.Println("=====")
}

func main() {
	teacher := Teacher{Id:34, Name:"Thor",Age:34}
	teacher.Show("Hawkeye")

	//下面通过反射，调用show方法
	v := reflect.ValueOf(teacher)
	//获取show方法
	m := v.MethodByName("Show")

	//校验一下，是否获取到show方法呢
	if !m.IsValid() {
		fmt.Printf("=======没有获取到制定的方法====")
		return
	}

	//参数必须是切片类型
	//reflect.Value{}   这里面，可以设定多个参数类型
	//目前，我们这里只有一个string类型的参数
	//
	args := []reflect.Value{reflect.ValueOf("Hulk")}
	m.Call(args)
}






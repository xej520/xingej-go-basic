// 相同结构体  之间 操作，如赋值，比较
package main

import "fmt"

type teacherA struct {
	string
	int
}
 type teacherB struct {
 	string
 	int
 }

func main() {

	boyTeacherA := teacherA{"xiaoli",22}
	//boyTeacherB := teacherB{"xiaoli",22}
	//説明：编译报错了，teacherA, teacherB  类型不相同，不能进行比较的
	//fmt.Println(boyTeacherA == boyTeacherB)

	boyTeacherB := teacherA{"xiaoli", 23}
	fmt.Println(boyTeacherB == boyTeacherA)

}






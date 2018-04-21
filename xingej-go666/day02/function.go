//函数  练习 测试
package main

func main() {

}

//==========函数声明=======
//==========方式一=====无输入参数======
//特点： 无输入参数，无返回值
func test() {

}
//特点： 无输入参数，有一个int返回值, 返回值类型，添加了 小括号
func test002() (int){

	return 0
}

//特点： 无输入参数，有一个int返回值, 返回值类型，也可以不添加小括号
//如果有两个以上的返回值时，必须添加括号
func test002a() int{

	return 0
}

//特点： 无输入参数，有两个返回值int，string, 返回值类型，也可以不添加小括号
func test002b() (int, string){

	return 0, "spark"
}

//==========方式2=====有输入参数======形式===
//特点：有参数
func test003(nodeName string, clusterName string , deployment string) {

}

// 特点：当前面参数的类型都一样时，可以只在相同类型参数最后一个添加类型，如下面形式
func test003a(nodeName , clusterName  , deployment string, cpu float64) {

}

//没有给返回值 命名
func test003b(nodeName , clusterName  , deployment string, cpu float64) (string, string, string) {

	return "spark", "k8s","ftp"
}

//给返回值 命名
func test003c(nodeName , clusterName  , deployment string, cpu float64) (a string, b string, c string) {

	//a,b,c 就是局部变量
	a, b, c = "spark", "k8s","ftp"

	return a, b, c
}

// 同样，当多个返回值的类型，相同时，只给最后一个添加类型，也是可以的
func test003d(nodeName , clusterName  , deployment string, cpu float64) (a , b , c string) {

	//a,b,c 就是局部变量
	a, b, c = "spark", "k8s","ftp"

	return a, b, c
}

//==========方式3=====变长参数======形式==
//特点：使用 三个点. 来表示
//要求：变长参数，必须是最后一个参数
// 传入的变长参数，实际上，是一个 切片类型哦
func test004(arg ... string) {

}

//要求：这种形式的用法，是不符合要求的，变长参数三个点，必须是最后一个参数
//func test004a(arg ... string, cpu float64) {
//
//}

//将上面的形式，改成下面的形式，就可以了
func test004b(cpu float64, arg ... string)  {

}








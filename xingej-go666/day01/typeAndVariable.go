package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c int = 1, 3, 4
	fmt.Println("a:\t",a,"\nb:\t",b,"\nc：\t",c)

	name,school,address := "laoxing","bjtu","beijing"
	fmt.Println("name：\t" ,name, "\nschool:\t",school,"\naddress:\t",address)

	var cha int = 65

	//将int类型的65，转换成字符类型的65
	chb := strconv.Itoa(cha)

	fmt.Println(chb)
	fmt.Println()

}

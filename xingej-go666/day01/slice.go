//切片slice  练习
package main

import "fmt"

func main() {
	//============先创建一个数组====
	// 创建数组，必须指明个数
	age := [4]int{}
	fmt.Println(age)

	//=========声明一个切片Slice类型=====
	// 看见了吧，切片类型，是不需要指明 个数的
	// 记住，slice类型的底层，也是数组
	var s1 []float64
	fmt.Println(s1)

	//========方式一：利用已有的数组，来初始化s1
	sparkNodeCpu := [...]float64{4.5, 3.4, 2.3, 9.8, 10}
	s1 = sparkNodeCpu[3:5] //从下标为3的元素，开始获取
	fmt.Println("------------->\t", s1)
	s1 = sparkNodeCpu[1:] //从下标为1的元素，获取到 最后
	fmt.Println(s1)

	s1 = sparkNodeCpu[:4] //指定最后一个获取的元素
	fmt.Println(s1)

	//========方式二：利用make关键字，来创建切片
	//make([]T, len, cap)
	//[]T,表示切片的类型
	//len,表示当前元素的个数，或者，初始化的个数
	//cap，由于切片的底层其实是数组，而数组在内存里是一块连续的空间，
	//		如make([]int, 3,10)  为了提升效率，一般先在内存里创建
	// 一块连续的内存大小，为10，如果你的元素个数超过了10的话，
	//Go语言，会默认将你现在的内存空间，由10，增加到20，重新在内存里找一块连续的空间，分配20个空间
	// 如果又超过20的话，就分配40，就这么下去。因此，最好你知道你需要多大的内存空间
	s2 := make([]int, 3, 10)
	fmt.Println(len(s2), cap(s2)) // 3  10

	s2a := make([]int, 3)           // 也可以不指定cap,这样的话，cap的默认值，就是len的长度
	fmt.Println(len(s2a), cap(s2a)) // 3  3

	//=======================Reslice========练习=====
	// 先声明一个切片类型
	a := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'j', 'k', 'u'}

	//以原切片a为基础，开始切
	//此时，你要注意了，新产生的切片s3的容量
	//下面的操作就是Reslice
	s3 := a[2:5] //表示，从下标为2开始，到5结束，最大不能超过被切slice的cap容量7
	s3a := s3[3:5]

	fmt.Println("长度:\t", len(s3), "\t容量cap:\t", cap(s3))   //3 7
	fmt.Println("长度:\t", len(s3a), "\t容量cap:\t", cap(s3a)) //2 4
	fmt.Println(string(s3a))

	//=======================Append========练习=====
	slice1 := make([]int, 3, 6)                   //默认值是全是0
	fmt.Printf("先打印出slice1的内存地址:\t %p\n", slice1) // 0xc042078090

	//先追加3个元素
	slice1 = append(slice1, 1, 3, 5)

	//再打印出slice1的值，和 内存地址
	fmt.Printf("%v  %p\n", slice1, slice1) //[0 0 0 1 3 5]  0xc042078090

	//继续追加元素，查看，内存地址，是否发生了变化
	slice1 = append(slice1, 4, 5, 6)

	// 超过了slice1的容量后，重新分配了内存地址
	fmt.Printf("%v  %p\n", slice1, slice1) //[0 0 0 1 3 5 4 5 6]  0xc042056060

	fmt.Println("=======================================")
	//测试，当多个切片都指向同一个底层数组时，并且多个切片有共同的元素时，如果其中一个元素，发生变化的话，
	//其他切片也会发生变化的
	appleSlice := []int{3, 4, 5, 9, 8, 7}
	bananaS := appleSlice[3:6]
	orangeS := appleSlice[2:5]

	//bananaS  orangeS  有两个共同的元素
	fmt.Println(appleSlice) //[3 4 5 9 8 7]
	fmt.Println(bananaS)    //[9 8 7]
	fmt.Println(orangeS)    //[5 9 8]

	fmt.Println("=======================================")
	//假设bananaS切片里的元素，第1个元素，发生了变化的话，
	//测试，appleSlice，orangeS 是否也发生了变呢
	//我改动的是下标为1，值为8的元素，将8改为了110，与此同时，其他切片中，8的值也全都改成了110了
	bananaS[1] = 110        //将下标为1的元素，设置为110
	fmt.Println(appleSlice) //[3 4 5 9 110 7]
	fmt.Println(bananaS)    //[9 110 7]
	fmt.Println(orangeS)    //[5 9 110]

	fmt.Println("=======================================")
	// 注意，此时，orangeS，orangeS 依旧都执行底层的数组appleSlice
	//还要注意下面的情形
	//利用Append方法
	bananaS = append(
		bananaS, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7)
	//此时，bananaS已经不再指向appleSlice了，因为超过了appleSlice的容量

	//而是，指向了新的地址，此时，我们再修改bananaS
	bananaS[1] = 1110

	fmt.Println(appleSlice) //[3 4 5 9 110 7]
	//看见了把，仅仅是修改了自己的值，appleSlice，orangeS 并没有发生变化
	fmt.Println(bananaS) //[9 1110 7 7 7 7 7 7 7 7 7 7 7 7 7 7]
	fmt.Println(orangeS) //[5 9 110]

	fmt.Println("=======================================")
	//=======================Copy========练习=====
	// copy(A,B) 是说，将B里的元素，拷贝到A里
	ftpNum := []int{4, 6, 7, 2, 3, 8, 9}
	sftpNum := []int{1, 2, 3}
	//将sftp
	copy(ftpNum, sftpNum) //注意，copy之后，就会将原先的值该了
	fmt.Println(ftpNum)
	fmt.Println(sftpNum)

	//截取拷贝
	copy(sftpNum[1:2], ftpNum[2:3])
	fmt.Println(ftpNum)
	fmt.Println(sftpNum)

}

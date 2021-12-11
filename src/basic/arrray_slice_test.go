package main

import (
	"fmt"
	"testing"
)

func create() {
	// 第一种方法
	var _ [3]int = [3]int{1, 2, 3}

	// 第二种方法
	_ = [3]int{1, 2, 3}

	//第三种方式，不指定大小
	_ = [...]int{1, 2, 3}
}

func create2() {

	/*如果你觉得每次写 [3]int 有点麻烦，你可以为 [3]int 定义一个类型字面量，也就是别名类型。
	使用 type 关键字可以定义一个类型字面量，后面只要你想定义一个容器大小为3，元素类型为int的数组 ，都可以使用这个别名类型。
	*/
	type arr3 [3]int

	myarr := arr3{2: 3}
	fmt.Printf("%d 的类型是: %T\n", myarr, myarr)
	fmt.Println(myarr[0:3])
}

//切片
func TestSlice(t *testing.T) {

	//第一种方式
	myarr := [5]int{1, 2, 3, 4, 5}
	fmt.Println()
	fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr), cap(myarr))

	mysli1 := myarr[1:3]
	fmt.Printf("mysli1 的长度为：%d，容量为：%d\n", len(mysli1), cap(mysli1))
	fmt.Println(mysli1)

	mysli2 := myarr[1:3:4]
	fmt.Printf("mysli2 的长度为：%d，容量为：%d\n", len(mysli2), cap(mysli2))
	fmt.Println(mysli2)

	//第二种方式
	// 声明字符串切片
	var _ []string

	// 声明整型切片
	var _ []int

	// 声明一个空切片
	var _ = []int{}

	//第三种方式：使用make
	//make 函数的格式：make( []Type, size, cap )

	//没指定cap则默认为len值
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))

}

func TestSliceMethod(t *testing.T) {
	myarr := []int{100}
	fmt.Println(myarr)
	// 追加一个元素
	myarr = append(myarr, 2)
	// 追加多个元素
	myarr = append(myarr, 3, 4)
	// 追加一个切片, ... 表示解包，不能省略
	myarr = append(myarr, []int{7, 8}...)
	// 在第一个位置插入元素
	myarr = append([]int{0}, myarr...)
	// 在中间插入一个切片(两个元素)
	myarr = append(myarr[:5], append([]int{5, 6}, myarr[5:]...)...)
	fmt.Println(myarr)
}

func TestSlice_2(t *testing.T) {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Println(myslice)
	fmt.Println(cap(myslice))

	myslice = myslice[:cap(myslice)]
	fmt.Println(myslice)
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}

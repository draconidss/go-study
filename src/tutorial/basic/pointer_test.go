package main

import (
	"fmt"
	"testing"
)

func TestCreation(t *testing.T) {
	//第一种
	// 定义普通变量
	aint := 1
	// 定义指针变量
	_ = &aint

	//第二种
	// 创建指针
	astr := new(string)
	// 给指针赋值
	*astr = "Go编程时光"

	//第三种
	aint = 1
	var _ *int // 声明一个指针=
	_ = &aint  // 初始化
}

func TestSymbol(t *testing.T) {
	aint := 1    // 定义普通变量
	ptr := &aint // 定义指针变量
	fmt.Println("普通变量存储的是：", aint)
	fmt.Println("普通变量存储的是：", *ptr)
	fmt.Println("指针变量存储的是：", &aint)
	fmt.Println("指针变量存储的是：", ptr)
}

func TestType(t *testing.T) {
	astr := "hello"
	aint := 1
	abool := false
	//arune := 'a'
	afloat := 1.2

	fmt.Printf("astr 指针类型是：%T\n", &astr)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)
}

func modifyArrayBySlice(sls []int) {
	sls[0] = 90
}
func modifyArrayByPtr(arr *[3]int) {
	(*arr)[1] = 91
}
func TestArrayAndSlice(t *testing.T) {
	a := [3]int{89, 90, 91}
	modifyArrayBySlice(a[:])
	modifyArrayByPtr(&a)
	fmt.Println(a)
}

func incr(p *int) int {
	a := *p
	a++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return a
}
func TestPointAsArgument(t *testing.T) {
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
	fmt.Println(v)
}

package main

import (
	"fmt"
	"testing"
)

//ref: https://golang.iswbm.com/c01/c01_02.html

// 1.使用var声明,var <name> <type>
func TestCreate_1(t *testing.T) {
	var par1 string
	var par2 = "字符串"
	var par3 int
	var par4 float32 = 10.1
	//默认为float64类型
	var par4_1 = 10.1
	var par5 bool

	fmt.Println("1.使用var声明,var <name> <type>")
	fmt.Println(par1)   // ""
	fmt.Println(par2)   // "字符串"
	fmt.Println(par3)   // 0
	fmt.Println(par4)   // 10.1
	fmt.Println(par4_1) // 10.1
	fmt.Println(par5)   // false
}

// 2.多个变量一起声明
func TestCreate_2(t *testing.T) {
	var (
		par6 int
		par7 float64
		par8 string
	)

	fmt.Println("2. 多个变量一起声明")
	fmt.Println(par6) // 0
	fmt.Println(par7) // 0
	fmt.Println(par8) // ""
}

// 3.使用 := 简短声明来显示初始化，只能在函数内部
func TestCreate_3(t *testing.T) {
	par9 := ":=初始化"
	fmt.Println("3.使用 := 来显示初始化，只能在函数内部")
	fmt.Println(par9)
}

// 4.声明和初始化多个变量
func TestCreate_Multiple(t *testing.T) {

	var (
		par10 float32
		par11 float32
	)

	//也常用于变量交换
	par10, par11 = par11, par10

	fmt.Println("4.声明和初始化多个变量")
	fmt.Println(par10)

	fmt.Println(par11)

}

// 5.使用表达式 new(Type) 将创建一个Type类型的匿名变量，初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type。
func TestCreate_New(t *testing.T) {

	//表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值
	//然后返回变量地址，返回的指针类型为*T
	par12 := new(int)

	fmt.Println("通过new的形式声明变量")
	fmt.Println(par12)
	fmt.Println(*par12)
	fmt.Println(&par12)
}

// 变量的零值
func TestDefaultValue(t *testing.T) {
	var (
		i int
		f float32
		b bool
		s string
	)
	// 0 0 false ""
	fmt.Printf("%v %v %v %q", i, f, b, s)
}

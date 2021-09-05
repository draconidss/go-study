package main

import "fmt"


//1.使用var声明,var <name> <type>
func create_1()  {
	var par1 string
	var par2 = "字符串"
	var par3 int
	var par4 float32 = 10.1
	//默认为float64类型
	var par4_1 = 10.1
	var par5 bool

	fmt.Println("1.使用var声明,var <name> <type>")
	fmt.Println(par1)
	fmt.Println(par2)
	fmt.Println(par3)
	fmt.Println(par4)
	fmt.Println(par4_1)
	fmt.Println(par5)
}


//2.多个变量一起声明
func create_2()  {
	var (
		par6 int
		par7 float64
		par8 string
	)

	fmt.Println("2. 多个变量一起声明")
	fmt.Println(par6)
	fmt.Println(par7)
	fmt.Println(par8)
}


//3.使用 := 来显示初始化，只能在函数内部
func create_3()  {
	par9 := ":=初始化"
	fmt.Println("3.使用 := 来显示初始化，只能在函数内部")
	fmt.Println(par9)
}


//4.声明和初始化多个变量
func create_4()  {


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


//5.使用表达式 new(Type) 将创建一个Type类型的匿名变量，初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type。
func create_5()  {

	//这里返回的是内存地址值,返回的是
	par12 := new(int)

	fmt.Println("通过new的形式声明变量")
	fmt.Println(par12)
	fmt.Println(*par12)
	fmt.Println(&par12)
}


func main() {
	create_1()
	create_2()
	create_3()
	create_4()
	create_5()
}
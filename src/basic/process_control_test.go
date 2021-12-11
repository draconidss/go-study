package main

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	age := 20
	if age > 18 {
		fmt.Println("已经成年了")
	} else if age > 12 {
		fmt.Println("已经是青少年了")
	} else {
		fmt.Println("还不是青少年")
	}

	if age := 20; age > 18 {
		fmt.Println("已经成年了")
	}

	gender := "male"
	if age > 18 && gender == "male" {
		fmt.Println("是成年男性")
	}
}

func TestSwitch(t *testing.T) {
	month := 2

	//多条件
	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误...")
	}

	score := 30

	//等同于if else
	switch {
	case score >= 95 && score <= 100:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("合格")
	case score >= 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误...")
	}

	//穿透，只能穿透一层，且不用判断下一层的条件
	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s == "xxxx":
		fmt.Println("xxxx")
	case s != "world":
		fmt.Println("world")
	}
}

func TestLoop(t *testing.T) {
	a := 1
	for a <= 5 {
		fmt.Println(a)
		a++
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//死循环类似while
	var i int = 1
	for {
		if i > 5 {
			break
		}
		fmt.Printf("hello, %d\n", i)
		i++
	}

	//for range,range 后可接数组、切片，字符串等
	myarr := [...]string{"java", "python", "golang"}
	for index, item := range myarr {
		fmt.Printf("%d:%s\n", index, item)
	}

	//如果你用一个变量来接收的话，接收到的是索引
	for index := range myarr {
		fmt.Printf("hello, %v\n", index)
	}

}

//goto 可以打破原有代码执行顺序，直接跳转到某一行执行代码。
func TestGoTo(t *testing.T) {
	//A
	goto flag
	fmt.Println("B")
flag:
	fmt.Println("A")

	//构建循环
	i := 1
loop:
	if i <= 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//退出循环,实现break
	j := 1
	for {
		if j > 5 {
			goto out
		}
		fmt.Println(j)
		j++
	}
out:

	//实现continue
	i = 1
keep:
	for i <= 10 {
		if i%2 == 1 {
			i++
			goto keep
		}
		fmt.Println(i)
		i++
	}

	//goto语句与标签之间不能有变量声明，否则编译错误。
	fmt.Println("start")
	goto flag_2
	//var say = "hello oldboy"
	fmt.Println()
flag_2:
	fmt.Println("end")
}

func deferFunc() {
	fmt.Println("B")
}

//函数的调用延迟到当前函数执行完后再执行。
func TestDefer(t *testing.T) {
	//deferFunc将在输出A后执行
	defer deferFunc()
	fmt.Println("A")

	//变量快照
	name := "go"
	defer fmt.Println(name) // 输出: go

	name = "python"
	fmt.Println(name) // 输出: python

	//匿名函数
	name = "go"
	defer func() {
		fmt.Println(name) // 输出: python
	}()
	name = "python"
	fmt.Println(name) // 输出: python
}

//反序调用,类似栈
func TestMultiDefer(t *testing.T) {
	name := "go"
	defer fmt.Println(name) // 输出: go

	name = "python"
	defer fmt.Println(name) // 输出: python

	name = "java"
	fmt.Println(name)
}

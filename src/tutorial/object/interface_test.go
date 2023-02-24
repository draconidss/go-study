package object

import (
	"fmt"
	"testing"
)

func TestUsage(t *testing.T) {
	//空接口的type和value都为nil
	var i interface{}
	fmt.Printf("type: %T, value: %v", i, i)

	//可以承载任意类型的值，也可以用空接口作为参数接收任意值
	// 存 int 没有问题
	i = 1
	fmt.Println(i)
	// 存字符串也没有问题
	i = "hello"
	fmt.Println(i)
	// 存布尔值也没有问题
	i = false
	fmt.Println(i)

	//定义一个接收任意类型的切片
	any := make([]interface{}, 5)
	any[0] = 11
	any[1] = "hello world"
	any[2] = []int{11, 22, 33, 44}
	for _, value := range any {
		fmt.Println(value)
	}
}

//注意点
func TestInterfaceTip(t *testing.T) {
	//1.接口无法向下转型即使转成空接口和转出空接口的类型都是一样
	// 声明a变量, 类型int, 初始值为1
	var a int = 1
	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	var i interface{} = a
	// 声明b变量, 尝试赋值i
	var b int = i
	fmt.Println(b)

	//2.当空接口承载数组和切片后，该对象无法再进行切片
	sli := []int{2, 3, 5, 7, 11, 13}
	var i_1 interface{} = sli
	g := i_1[1:3]
	fmt.Println(g)

	//3.当你使用空接口来接收任意类型的参数时，它的静态类型是 interface{}，
	//但动态类型（是 int，string 还是其他类型）我们并不知道，因此需要使用类型断言。

}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) send_wechat() {
	fmt.Println("Hello, Wechat.")
}

//测试实现接口的结构体方法的调用限制
func TestMethodExtendError(t *testing.T) {
	//这里声明类型为定义的Phone接口
	var phone Phone
	phone = iPhone{name: "ming's iphone"}
	phone.call()
	//无法调用该方法
	phone.send_wechat()
}

//修改为不显示的声明为Phone接口类型
func TestMethodExtendSuccess(t *testing.T) {
	phone := iPhone{name: "ming's iphone"}
	phone.call()
	phone.send_wechat()
}

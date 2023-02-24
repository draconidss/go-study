package object

import (
	"fmt"
	"testing"
)

//Type Assertion（中文名叫：类型断言）
//1.仅能对静态类型为空接口（interface{}）的对象进行断言
//2.类型断言完成后，实际上会返回静态类型为你断言的类型的对象，而要清楚原来的静态类型为空接口类型（interface{}）

//通过它可以做到以下几件事情
//检查 i 是否为 nil
//检查 i 存储的值是否为某个类型

//第一种：t := i.(T)
func TestAssertion_1(t *testing.T) {
	//var i interface{} = nil
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	fmt.Println("=====分隔线=====")

	t2 := i.(string)
	//触发panic
	fmt.Println(t2)
}

//第二种：t, ok:= i.(T)
//如果断言失败不会触发panic，而是将ok设为false，此时t为T的零值
func TestTestAssertion_2(t *testing.T) {
	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	t2, ok := i.(float64)
	fmt.Printf("%f-%t\n", t2, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("=====分隔线3=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)
}

//区分多种类型,这里会隐式转型为接口类型
func findType(i interface{}) {
	//只能在switch中使用
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

//主动转换为接口类型
func findTypeByInterface(i interface{}) {
	//主动转换
	findType(interface{}(i))
}

func TestTypeSwitch(t *testing.T) {
	findTypeByInterface(10)      // int
	findTypeByInterface("hello") // string

	var k interface{} // nil
	findTypeByInterface(k)

	findTypeByInterface(10.23) //float64
}

package object

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectThreeRules(t *testing.T) {
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	fmt.Println("第一定律：反射可以将接口类型变量 转换为“反射类型对象”")
	parmType := reflect.TypeOf(age)
	parmValue := reflect.ValueOf(age)

	//第一定律：反射可以将接口类型变量 转换为“反射类型对象”
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", parmType)   //*reflect.rtype
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", parmValue) //reflect.Value

	//第二定律：反射可以将 “反射类型对象”转换为接口类型变量；
	//最后转换后的对象，静态类型为 interface{} ，如果要转成最初的原始类型，需要再类型断言转换一下
	fmt.Println("第二定律：反射可以将 “反射类型对象”转换为 接口类型变量；")
	i := parmValue.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i) //从反射对象到接口变量：新对象的类型为 int 值为 25
	//通过断言转回来
	i = i.(int)

	//第三定律：如果要修改 “反射类型对象” 其类型必须是可写的
	fmt.Println("第三定律：如果要修改 “反射类型对象” 其类型必须是可写的")
	var name string = "Go编程时光"
	v1 := reflect.ValueOf(name)
	fmt.Println("v1 可写性为:", v1.CanSet()) //false

	//传入变量指针并使用Elem()函数返回指针指向的数据
	v2 := reflect.ValueOf(&name).Elem()
	fmt.Println("v2 可写性为:", v2.CanSet()) //true
	v2.SetString("修改后的值")
	fmt.Println("通过反射对象进行更新后，真实世界里 name 变为：", name)

}

type ReflectStruct struct {
	name   string
	age    int
	gender string
}

//获取类别
func TestReflectGetKind(t *testing.T) {
	m := ReflectStruct{}

	//第一种：传入值
	fmt.Println("第一种：传入值")
	parmType := reflect.TypeOf(m)
	fmt.Println("Type: ", parmType)        //object.ReflectStruct
	fmt.Println("Kind: ", parmType.Kind()) //struct

	//第二种：传入指针
	fmt.Println("第二种：传入指针")
	parmType = reflect.TypeOf(&m)
	fmt.Println("&m Type: ", parmType)              //&m Type:  *object.ReflectStruct
	fmt.Println("&m Kind: ", parmType.Kind())       //&m Kind:  ptr
	fmt.Println("m Type: ", parmType.Elem())        //m Type:  object.ReflectStruct
	fmt.Println("m Kind: ", parmType.Elem().Kind()) //m Kind:  struct

}

//数据转换
func TestReflectDataSwitch(t *testing.T) {
	var age int = 25

	v1 := reflect.ValueOf(age)
	fmt.Printf("转换int前， type: %T, value: %v \n", v1, v1) //转换前， type: reflect.Value, value: 25
	v2 := v1.Int()
	fmt.Printf("转换int后， type: %T, value: %v \n", v2, v2) //转换后， type: int64, value: 25

	v1 = reflect.ValueOf(&age)
	fmt.Printf("转换指针前， type: %T, value: %v \n", v1, v1)
	v3 := v1.Pointer()
	fmt.Printf("转换指针后， type: %T, value: %v \n", v3, v3)
}

//对切片操作
func TestReflectSlice(t *testing.T) {
	var numList []int = []int{1, 2}

	v1 := reflect.ValueOf(numList)
	fmt.Printf("转换前，type：%T, value: %v \n", v1, v1) //转换前，type：reflect.Value, value: [1 2]

	// Slice 函数接收两个参数
	v2 := v1.Slice(0, 1)
	//返回还是 reflect.Value
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2) //转换后， type: reflect.Value, value: [1]

	//更新切片值
	fmt.Println("更新切片值")
	arr := []int{1, 2}
	value := reflect.ValueOf(&arr).Elem()
	value.Set(reflect.Append(value, reflect.ValueOf(3)))
	fmt.Println(value)
	fmt.Println(value.Len())
}

type Reflect struct {
	name string
	age  int
}

//对属性的操作
func TestReflectField(t *testing.T) {
	p := Reflect{name: "Draco", age: 18}

	value := reflect.ValueOf(p)

	fmt.Println("字段数:", value.NumField())
	fmt.Println("第 1 个字段的值：", value.Field(0))
	fmt.Println("第 2 个字段的值：", value.Field(1))

	fmt.Println("==========================")
	// 也可以这样来遍历
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("第 %d 个字段：%v \n", i+1, value.Field(i))
	}
}

func (r Reflect) SayBye() string {
	return "Bye"
}

func (r Reflect) SayHello() string {
	return "Hello"
}

func (r Reflect) SelfIntroduction(name string, age int) string {
	return "Hello, my name is " + name + "and im " + string(rune(age)) + "years old."
}

//对方法的操作
//对方法本身用TypeOf
//动态调用方法用ValueOf
func TestReflectMethod(t *testing.T) {
	p := Reflect{name: "Draco", age: 18}

	parmType := reflect.TypeOf(&p)
	parmValue := reflect.ValueOf(&p)

	fmt.Println("方法数（可导出的）:", parmType.NumMethod())
	fmt.Println("第 1 个方法：", parmType.Method(0).Name)
	fmt.Println("第 2 个方法：", parmType.Method(1).Name)
	fmt.Println("==========================")
	// 也可以这样来遍历
	for i := 0; i < parmType.NumMethod()-1; i++ {
		fmt.Printf("调用第 %d 个方法：%v ，调用结果：%v\n",
			i+1,
			parmType.Method(i).Name,
			parmValue.MethodByName(parmType.Method(i).Name).Call(nil))
	}

	//调用有参函数
	fmt.Println("==========================")
	input := make([]reflect.Value, 2)
	input[0] = reflect.ValueOf("Colin")
	input[1] = reflect.ValueOf(54)
	result := parmValue.MethodByName("SelfIntroduction").Call(input)
	fmt.Println(result)
	fmt.Println(result[0].Interface())
}

//动态调用函数（使用索引且无参数）

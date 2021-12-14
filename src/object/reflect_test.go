package object

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	parmType := reflect.TypeOf(age)
	parmValue := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", parmType)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", parmValue)
}

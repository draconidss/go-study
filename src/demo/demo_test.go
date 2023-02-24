package main

import (
	"fmt"
	"reflect"
	"testing"
)

//reflect 检查类型
func TestCheckTypeByReflect(t *testing.T) {
	var (
		s string = "string"
		//默认为 int
		i = 10
		//默认为 float64，尽量手动声明为 float32
		f = 1.2
	)

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}

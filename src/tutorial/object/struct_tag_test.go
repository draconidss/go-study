package object

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

//--------------------------------------标签tag------------------------------------------
//字段上还可以额外再加一个属性，用反引号（Esc键下面的那个键）包含的字符串，称之为 Tag，也就是标签。
type Person struct {
	Name string `json:"name" label:"Name is: "`
	//只要发现被omitempty标记的属性 为 false， 0， 空指针，空接口，空数组，空切片，空映射，空字符串中的一种，就会被忽略。
	Age  int    `json:"age,omitempty" label:"Age is: "`
	Addr string `json:"addr,omitempty" label:"Gender is: " default:"unknown"`
}

func TestTag(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  0,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1 没有 Addr且age为0，就不会打印了Addr和age属性
	fmt.Printf("%s\n", data1)

	// ================

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)
}

func TestGetTagByReflect(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	// 三种获取 field
	field_1, _ := reflect.TypeOf(p1).FieldByName("Age")
	//field_2 := reflect.ValueOf(p1).Type().Field(0)         // i 表示第几个字段
	//field_3 := reflect.ValueOf(&p1).Elem().Type().Field(1) // i 表示第几个字段

	tag := field_1.Tag
	fmt.Println(tag)

	// 获取键值对
	//Get是对Lookup的封装
	labelValue := tag.Get("json")
	labelValue_1, _ := tag.Lookup("json")

	fmt.Println(labelValue)
	fmt.Println(labelValue_1)

}

func Print(obj interface{}) error {
	// 取 Value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}

//通过tag操作实现打印：
//Name is: MING
//Age is: 29
//Gender is: unknown
func TestPrint(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "",
	}
	Print(p1)
}
